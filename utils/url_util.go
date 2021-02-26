package utils

import (
	"fmt"

	"github.com/red-gold/telar-core/config"
)

// GetPrettyURL return *config.AppConfig.BaseRoute
func GetPrettyURL() string {
	return *config.AppConfig.BaseRoute
}

// GetPrettyURL formats according to pretty URL from (baseFunctionURL+url) and returns the resulting string.
func GetPrettyURLf(url string) string {
	fmt.Println("FROM core *config.AppConfig.BaseRoute ", *config.AppConfig.BaseRoute)
	return fmt.Sprintf("%s%s", *config.AppConfig.BaseRoute, url)
}
