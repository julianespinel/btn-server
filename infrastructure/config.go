package infrastructure

type DBConfig struct {
	Username string
	Password string
	DbName   string
}

type Config struct {
	DbConfig DBConfig
}
