package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func loadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatal("Load Env failed. %v", err)
	}
	return &config{
		app: &app,
		db:  &db,
		jwt: &jwt,
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
}
type IDbConfig interface {
}
type IJwtConfig interface {
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	adminKey     string
	fileLimit    int
	gcpbucket    string
}

func (c *config) App() IAppConfig {
	return nil
}

type db struct {
	host          string
	port          int
	protocol      string
	username      string
	password      string
	database      string
	sslmode       string
	maxConnection int
}

func (c *config) Db() IDbConfig {
	return nil
}

type jwt struct {
	adminKey        string
	secertKey       string
	apikey          string
	accessExpiresAt int
	refeshExpiresAt int
}

func (c *config) Jwt() IJwtConfig {
	return nil
}
