package authcookie

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/red-gold/telar-core/pkg/log"
	"github.com/red-gold/telar-core/pkg/parser"
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

		// Get authorization cookie
		auth := readAuthCookie(c, cfg.HeaderCookieName, cfg.PayloadCookieName, cfg.SignatureCookieName)

		// Check if the jwt token contains content
		jwtToken := strings.Split(auth, ".")
		if len(jwtToken[0]) == 0 || len(jwtToken[1]) == 0 || len(jwtToken[2]) == 0 {
			log.Error("Token does not contains content %s ", auth)
			return cfg.Unauthorized(c)
		}

		// Check if the JWT secret key is not nill
		if cfg.JWTSecretKey == nil {
			log.Error("JWT secret key is not provided in config!")
			return c.SendStatus(http.StatusInternalServerError)
		}

		// Check token validation and set user context in locals
		if parsedClaim, err := cfg.Authorizer(auth); err == nil && parsedClaim != nil {

			userCtx := new(types.UserContext)
			parser.MarshalMap(parsedClaim["claim"], userCtx)

			c.Locals(cfg.UserCtxName, *userCtx)
			return c.Next()

		} else {
			log.Error("Unuthorize user %s ", err.Error())
		}

		// Authentication failed
		return cfg.Unauthorized(c)
	}
}
