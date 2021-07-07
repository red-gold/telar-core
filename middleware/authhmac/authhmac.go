package authhmac

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
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

		// Get authorization header
		auth := c.Get(types.HeaderHMACAuthenticate)

		// Check if the HMAC header contains content
		if len(auth) == 0 {
			log.Error("Unauthorized! HMAC not presented!")
			return cfg.Unauthorized(c)
		}

		if err := cfg.Authorizer(c.Body(), auth); err == nil {
			if c.Get("uid") == "" {
				log.Warn("[HMAC] User id is not provided. In this case user context will be set empty!")
				c.Locals(cfg.UserCtxName, types.UserContext{})
				return c.Next()
			}
			userUUID, userUuidErr := uuid.FromString(c.Get("uid"))
			if userUuidErr == nil {
				c.Locals(cfg.UserCtxName, types.UserContext{
					UserID:      userUUID,
					Username:    c.Get("email"),
					Avatar:      c.Get("avatar"),
					SocialName:  c.Get("socialName"),
					DisplayName: c.Get("displayName"),
					SystemRole:  c.Get("role"),
				})
				return c.Next()
			} else {
				log.Error("Can not parse UID from claim %s", userUuidErr.Error())
			}
		} else {
			log.Error("HMAC validation %s", err.Error())
		}

		// Authentication failed
		return cfg.Unauthorized(c)
	}
}
