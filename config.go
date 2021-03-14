package core

import (
	"os"
	"strconv"

	"github.com/red-gold/telar-core/config"
	"github.com/red-gold/telar-core/pkg/log"
)

// Initialize AppConfig
func InitConfig() {

	// Load config from environment values if exists
	LoadConfigFromEnvironment()
}

// Initialize AppConfig
func InitConfigFromData(newConfig config.Configuration) {
	config.AppConfig = newConfig
}

// Load config from environment
func LoadConfigFromEnvironment() {

	appName, ok := os.LookupEnv("app_name")
	if ok {
		config.AppConfig.AppName = &appName
		log.Info("App Name information loaded from env.")
	}

	debug, ok := os.LookupEnv("debug")
	if ok {
		parsedDebug, errParseDebug := strconv.ParseBool(debug)
		if errParseDebug != nil {
			log.Error("[ERROR]: Debug information loading error: %s", errParseDebug.Error())
		}
		config.AppConfig.Debug = &parsedDebug
		log.Info("Debug information loaded from env.")
	}
	gateway, ok := os.LookupEnv("gateway")
	if ok {
		config.AppConfig.Gateway = &gateway
		log.Info("Gateway information loaded from env.")
	}

	internalGateway, ok := os.LookupEnv("internal_gateway")
	if ok {
		config.AppConfig.InternalGateway = &internalGateway
		log.Info("Internal gateway information loaded from env. | %s |", internalGateway)
	}

	webDomain, ok := os.LookupEnv("web_domain")
	if ok {
		config.AppConfig.WebDomain = &webDomain
		log.Info("Web domain information loaded from env.")
	}

	orgName, ok := os.LookupEnv("org_name")
	if ok {
		config.AppConfig.OrgName = &orgName
		log.Info("Organization Name information loaded from env.")
	}

	orgAvatar, ok := os.LookupEnv("org_avatar")
	if ok {
		config.AppConfig.OrgAvatar = &orgAvatar
		log.Info("Organization Avatar information loaded from env.")
	}

	server, ok := os.LookupEnv("server")
	if ok {
		config.AppConfig.Server = &server
		log.Info("Server information loaded from env.")
	}

	payloadSecret, ok := os.LookupEnv("payload_secret")
	if ok {
		config.AppConfig.PayloadSecret = &payloadSecret
		log.Info("Payload secret information loaded from env.")
	}

	publicKey, ok := os.LookupEnv("public_key")
	if ok {
		config.AppConfig.PublicKey = &publicKey
		log.Info("Public key information loaded from env.")
	}

	privateKey, ok := os.LookupEnv("private_key")
	if ok {
		config.AppConfig.PrivateKey = &privateKey
		log.Info("PrivateKey information loaded from env.")
	}

	recaptchaKey, ok := os.LookupEnv("recaptcha_key")
	if ok {
		config.AppConfig.RecaptchaKey = &recaptchaKey
		log.Info("Recaptcha key information loaded from env.")
	}

	recaptchaSiteKey, ok := os.LookupEnv("recaptcha_site_key")
	if ok {
		config.AppConfig.RecaptchaSiteKey = &recaptchaSiteKey
		log.Info("Recaptcha site key information loaded from env.")
	}

	origin, ok := os.LookupEnv("origin")
	if ok {
		config.AppConfig.Origin = &origin
		log.Info("Origin information loaded from env.")
	}

	headerCookieName, ok := os.LookupEnv("header_cookie_name")
	if ok {
		config.AppConfig.HeaderCookieName = &headerCookieName
		log.Info("Header cookie name information loaded from env.")
	}

	payloadCookieName, ok := os.LookupEnv("payload_cookie_name")
	if ok {
		config.AppConfig.PayloadCookieName = &payloadCookieName
		log.Info("Payload cookie name information loaded from env.")
	}

	signatureCookieName, ok := os.LookupEnv("signature_cookie_name")
	if ok {
		config.AppConfig.SignatureCookieName = &signatureCookieName
		log.Info("Signature cookie name information loaded from env.")
	}

	mongodbHost, ok := os.LookupEnv("mongo_host")
	if ok {
		config.AppConfig.MongoDBHost = &mongodbHost
		log.Info("MongoDB host information loaded from env.")
	}

	baseRoute, ok := os.LookupEnv("base_route")
	if ok {
		config.AppConfig.BaseRoute = &baseRoute
		log.Info("Base route information loaded from env.")
	}

	database, ok := os.LookupEnv("mongo_database")
	if ok {
		config.AppConfig.Database = &database
		log.Info("MongoDB database information loaded from env.")
	}

	smtpEmail, ok := os.LookupEnv("smtp_email")
	if ok {
		config.AppConfig.SmtpEmail = &smtpEmail
		log.Info("SMTP Email information loaded from env.")
	}

	refEmail, ok := os.LookupEnv("ref_email")
	if ok {
		config.AppConfig.RefEmail = &refEmail
		log.Info("Reference Email information loaded from env.")
	}

	phoneSourceNumebr, ok := os.LookupEnv("phone_source_number")
	if ok {
		config.AppConfig.PhoneSourceNumber = &phoneSourceNumebr
		log.Info("Phone Source Number information loaded from env.")
	}

	phoneAuthToken, ok := os.LookupEnv("phone_auth_token")
	if ok {
		config.AppConfig.PhoneAuthToken = &phoneAuthToken
		log.Info("Phone Auth Token  information loaded from env.")
	}

	phoneAuthId, ok := os.LookupEnv("phone_auth_id")
	if ok {
		config.AppConfig.PhoneAuthId = &phoneAuthId
		log.Info("Phone Auth Id  information loaded from env.")
	}

	refEmailPass, ok := os.LookupEnv("ref_email_pass")
	if ok {
		config.AppConfig.RefEmailPass = &refEmailPass
		log.Info("Reference Email Password  information loaded from env.")
	}

	dbType, ok := os.LookupEnv("db_type")
	if ok {
		config.AppConfig.DBType = &dbType
		log.Info("Database type information loaded from env.")
	}
}
