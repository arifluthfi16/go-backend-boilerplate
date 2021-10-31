package config

type ServerConf struct{
	DBConfig DBConf
	AppConfig AppConf
}

type DBConf struct {
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DBPass string
}

type AppConf struct {
	AppName string
	AppVersion string
}

var (
	ServerConfig = ServerConf{}
)


