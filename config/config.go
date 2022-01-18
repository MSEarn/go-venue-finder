package config

type Config struct {
	Server  Server
	Mysql   Mysql
	Logging Logging
}

type Mysql struct {
	Host            string
	Port            string
	Username        string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxConnLifetime int
}

type Logging struct {
	Level  string
	Format string
}

type Server struct {
	Port string
}
