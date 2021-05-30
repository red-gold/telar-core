package authrole

import (
	"github.com/gofiber/fiber/v2"
	"github.com/red-gold/telar-core/pkg/log"
	"github.com/red-gold/telar-core/types"
)

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Get current user
		currentUser, ok := c.Locals("user").(types.UserContext)
		if !ok {
			log.Error("Can not retrieve current user context")
			return cfg.Unauthorized(c)
		}

		if canAccess := cfg.Authorizer(currentUser.SystemRole); canAccess == true {
			return c.Next()
		} else {
			log.Error("Current user don't have enough privilege to access this route %s - %s ", currentUser.Username, c.Path())
		}

		// Authentication failed
		return cfg.Unauthorized(c)
	}
}
