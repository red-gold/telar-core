package authrole

import (
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
	// It will be called with a current user role
	// and is expected to return true or false to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(userRole string) bool

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	// Role control what resource can be available to different type of users
	//
	// Optional. Default: "admin"
	Role string

	// UserCtxName is the key to store the user context in Locals
	//
	// Optional. Default: "user"
	UserCtxName string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:         nil,
	Authorizer:   nil,
	Unauthorized: nil,
	Role:         "admin",
	UserCtxName:  types.UserCtxName,
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
		cfg.Authorizer = func(userRole string) bool {
			return cfg.Role == userRole
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			c.Set(fiber.HeaderWWWAuthenticate, "Role realm="+cfg.Realm)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	if cfg.Role == "" {
		cfg.Role = ConfigDefault.Role
	}
	if cfg.UserCtxName == "" {
		cfg.UserCtxName = ConfigDefault.UserCtxName
	}
	return cfg
}
