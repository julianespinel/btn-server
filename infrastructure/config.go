package infrastructure

type ServerConfig struct {
	Port int
}

type DBConfig struct {
	Username string
	Password string
	DbName   string
}

type SmsConfig struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

type Config struct {
	Server   ServerConfig
	Database DBConfig
	Sms      SmsConfig
}
