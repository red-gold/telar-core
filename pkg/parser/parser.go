package parser

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/gofrs/uuid"
	"github.com/gorilla/schema"
)

const queryTag = "query"

// MarshalMap
func MarshalMap(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}

// QueryParser binds the query string to a struct.
func QueryParser(c *fiber.Ctx, out interface{}) error {
	// Get decoder from pool
	var decoder = decoderPool.Get().(*schema.Decoder)
	defer decoderPool.Put(decoder)

	// Set correct alias tag
	decoder.SetAliasTag(queryTag)

	data := make(map[string][]string)
	c.Context().QueryArgs().VisitAll(func(key []byte, val []byte) {
		k := fiberUtils.UnsafeString(key)
		v := fiberUtils.UnsafeString(val)
		if strings.Contains(v, ",") && equalFieldType(out, reflect.Slice, k) {
			values := strings.Split(v, ",")
			for i := 0; i < len(values); i++ {
				data[k] = append(data[k], values[i])
			}
		} else {
			data[k] = append(data[k], v)
		}
	})

	return decoder.Decode(out, data)
}

// uuidConverter uuid convertor for schema decoder
var uuidConverter = func(value string) reflect.Value {
	if v, err := uuid.FromString(value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{} // this is the same as the private const invalidType
}

// decoderPool helps to improve BodyParser's and QueryParser's performance
var decoderPool = &sync.Pool{New: func() interface{} {
	var decoder = schema.NewDecoder()
	decoder.ZeroEmpty(true)
	decoder.RegisterConverter(uuid.UUID{}, uuidConverter)
	decoder.IgnoreUnknownKeys(true)
	return decoder
}}

func equalFieldType(out interface{}, kind reflect.Kind, key string) bool {
	// Get type of interface
	outTyp := reflect.TypeOf(out).Elem()
	key = fiberUtils.ToLower(key)
	// Must be a struct to match a field
	if outTyp.Kind() != reflect.Struct {
		return false
	}
	// Copy interface to an value to be used
	outVal := reflect.ValueOf(out).Elem()
	// Loop over each field
	for i := 0; i < outTyp.NumField(); i++ {
		// Get field value data
		structField := outVal.Field(i)
		// Can this field be changed?
		if !structField.CanSet() {
			continue
		}
		// Get field key data
		typeField := outTyp.Field(i)
		// Get type of field key
		structFieldKind := structField.Kind()
		// Does the field type equals input?
		if structFieldKind != kind {
			continue
		}
		// Get tag from field if exist
		inputFieldName := typeField.Tag.Get(queryTag)
		if inputFieldName == "" {
			inputFieldName = typeField.Name
		} else {
			inputFieldName = strings.Split(inputFieldName, ",")[0]
		}
		// Compare field/tag with provided key
		if fiberUtils.ToLower(inputFieldName) == key {
			return true
		}
	}
	return false
}
