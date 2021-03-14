package config

type (
	Configuration struct {
		BaseRoute           *string
		Server              *string
		Gateway             *string
		InternalGateway     *string
		MongoDBHost         *string
		Database            *string
		PayloadSecret       *string
		PublicKey           *string
		RecaptchaKey        *string
		RecaptchaSiteKey    *string
		HeaderCookieName    *string
		PayloadCookieName   *string
		SignatureCookieName *string
		SmtpEmail           *string
		RefEmail            *string
		RefEmailPass        *string
		Origin              *string
		PrivateKey          *string
		AppName             *string
		PhoneSourceNumber   *string
		PhoneAuthToken      *string
		PhoneAuthId         *string
		OrgName             *string
		OrgAvatar           *string
		WebDomain           *string
		DBType              *string
		QueryPrettyURL      *bool
		Debug               *bool
	}
)

const (
	DB_INMEMORY = "inmemory"
	DB_MONGO    = "mongo"
	DB_SQLITE   = "sqlite"
	DB_MYSQL    = "mysql"
)

// AppConfig holds the Configuration values from app-config.yml file
var AppConfig Configuration
