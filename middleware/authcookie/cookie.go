package authcookie

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// readAuthCookie Read cookie and get auth JWT string
func readAuthCookie(c *fiber.Ctx, HeaderCookieName, PayloadCookieName, SignatureCookieName string) string {
	return fmt.Sprintf("%s.%s.%s", c.Cookies(HeaderCookieName), c.Cookies(PayloadCookieName), c.Cookies(SignatureCookieName))
}

func authCookiePresented(c *fiber.Ctx, HeaderCookieName, PayloadCookieName, SignatureCookieName string) {

}
