package authhmac

import (
	"github.com/alexellis/hmac"
	"github.com/gofiber/fiber/v2"
	"github.com/red-gold/telar-core/types"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Realm is a string to define realm attribute of BasicAuth.
	// the realm identifies the system to authenticate against
	// and can be used by clients to save credentials
	//
	// Optional. Default: "Restricted".
	Realm string

	// Authorizer defines a function you can pass
	// to check the credentials however you want.
	// It will be called with a request body bytes(bytesIn), encoded hash
	// and is expected to return nil or error to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(bytesIn []byte, encodedHash string) error

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	// PayloadSecret is the key to validate HMAC
	//
	// Optional. Default: "secret"
	PayloadSecret string

	// UserCtxName is the key to store the user context in Locals
	//
	// Optional. Default: "user"
	UserCtxName string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:          nil,
	Authorizer:    nil,
	Unauthorized:  nil,
	PayloadSecret: "secret",
	UserCtxName:   types.UserCtxName,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.Authorizer == nil {
		cfg.Authorizer = func(bytesIn []byte, encodedHash string) error {
			return hmac.Validate(bytesIn, encodedHash, string(cfg.PayloadSecret))
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			c.Set(fiber.HeaderWWWAuthenticate, "HMAC realm="+cfg.Realm)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	if cfg.PayloadSecret == "" {
		cfg.PayloadSecret = ConfigDefault.PayloadSecret
	}
	if cfg.UserCtxName == "" {
		cfg.UserCtxName = ConfigDefault.UserCtxName
	}
	return cfg
}
