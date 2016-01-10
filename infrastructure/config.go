package infrastructure

type ServerConfig struct {
	Port int
}

type DBConfig struct {
	Username string
	Password string
	DbName   string
}

type Config struct {
	Server   ServerConfig
	Database DBConfig
}
