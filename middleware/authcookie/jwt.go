package authcookie

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// validateToken Validate and parse JWT token
func validateToken(key []byte, token string) (jwt.MapClaims, error) {
	publicKey, keyErr := jwt.ParseECPublicKeyFromPEM(key)
	if keyErr != nil {
		return nil, fmt.Errorf("Unable to read public key : %s", keyErr.Error())

	}

	parsed, parseErr := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if parseErr != nil {
		return nil, fmt.Errorf("Unable to decode token %s", parseErr.Error())

	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Token claim is not valid!")
	}
}
