package authcookie

import (
	"github.com/dgrijalva/jwt-go"
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
	// It will be called with a JWT token
	// and is expected to return user claim or error to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(token string) (jwt.MapClaims, error)

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	// JWTSecretKey is the key to validate JWT token
	//
	// Optional. Default: nil
	JWTSecretKey []byte

	// UserCtxName is the key to store the user context in Locals
	//
	// Optional. Default: "user"
	UserCtxName string

	// HeaderCookieName is the name of cookie keeping JWT header
	//
	// Optional. Default: "he"
	HeaderCookieName string

	// PayloadCookieName is the name of cookie keeping JWT payload
	//
	// Optional. Default: "pa"
	PayloadCookieName string

	// SignatureCookieName is the name of cookie keeping JWT signature
	//
	// Optional. Default: "si"
	SignatureCookieName string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:                nil,
	Authorizer:          nil,
	Unauthorized:        nil,
	JWTSecretKey:        nil,
	UserCtxName:         types.UserCtxName,
	HeaderCookieName:    "he",
	PayloadCookieName:   "pa",
	SignatureCookieName: "si",
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
		cfg.Authorizer = func(token string) (jwt.MapClaims, error) {
			return validateToken([]byte(cfg.JWTSecretKey), token)
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			c.Set(fiber.HeaderWWWAuthenticate, "Bearer realm="+cfg.Realm)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	if cfg.JWTSecretKey == nil {
		cfg.JWTSecretKey = ConfigDefault.JWTSecretKey
	}
	if cfg.UserCtxName == "" {
		cfg.UserCtxName = ConfigDefault.UserCtxName
	}
	if cfg.HeaderCookieName == "" {
		cfg.HeaderCookieName = ConfigDefault.HeaderCookieName
	}
	if cfg.PayloadCookieName == "" {
		cfg.PayloadCookieName = ConfigDefault.PayloadCookieName
	}
	if cfg.SignatureCookieName == "" {
		cfg.SignatureCookieName = ConfigDefault.SignatureCookieName
	}
	return cfg
}
