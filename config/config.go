package config

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func ConvertStrToNumber(val, msg string) int {
	p, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(msg)
	}
	return p
}
func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Load Env failed. %v", err)
	}
	return &config{
		app: &app{
			host:         envMap["APP_HOST"],
			port:         ConvertStrToNumber(envMap["APP_PORT"], "Load APP_PORT failed"),
			name:         envMap["APP_NAME"],
			version:      envMap["APP_VERSION"],
			bodyLimit:    ConvertStrToNumber(envMap["APP_BODY_LIMIT"], "Load APP_BODY_LIMIT failed"),
			readTimeout:  time.Duration(ConvertStrToNumber(envMap["APP_READ_TIMEOUT"], "Load APP_READ_TIMEOUT failed")),
			writeTimeout: time.Duration(ConvertStrToNumber(envMap["APP_WRITE_TIMEOUT"], "Load APP_WRITE_TIMEOUT failed")),
			fileLimit:    ConvertStrToNumber(envMap["APP_FILE_LIMIT"], "Load APP_FILE_LIMIT failed"),
			gcpbucket:    envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host:          envMap["DB_HOST"],
			port:          ConvertStrToNumber(envMap["DB_PORT"], "Load DB_PORT failed"),
			protocol:      envMap["DB_PROTOCOL"],
			username:      envMap["DB_USERNAME"],
			password:      envMap["DB_PASSWORD"],
			database:      envMap["DB_NAME"],
			sslmode:       envMap["DB_SSLMODE"],
			maxConnection: ConvertStrToNumber(envMap["DB_MAXCONNECTION"], "Load DB_MAXCONNECTION failed"),
		},
		jwt: &jwt{
			apikey:           envMap["JWT_API_KEY"],
			accessExpiresAt:  ConvertStrToNumber(envMap["JWT_ACESSS_EXPIRES"], "Load JWT_ACESSS_EXPIRES failed"),
			refreshExpiresAt: ConvertStrToNumber(envMap["JWT_REFESH_EXPIRES"], "Load JWT_REFRESH_EXPIRES failed"),
			adminKey:         envMap["JWT_ADMIN_KEY"],
			secretKey:        envMap["JWT_ACCESS_SECRET_KEY"],
		},
	}
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	fileLimit    int
	gcpbucket    string
}

type IAppConfig interface {
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GcpBucket() string
}

func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GcpBucket() string           { return a.gcpbucket }

func (c *config) App() IAppConfig {
	return c.app
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

type IDbConfig interface {
	Url() string
	MaxConnection() int
	Protocol() string
	Username() string
	Password() string
	Database() string
	SslMode() string
}

func (d *db) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", d.host, d.port, d.username, d.password, d.database, d.sslmode)
}
func (d *db) MaxConnection() int { return d.maxConnection }
func (d *db) Protocol() string   { return d.protocol }
func (d *db) Username() string   { return d.username }
func (d *db) Password() string   { return d.password }
func (d *db) Database() string   { return d.database }
func (d *db) SslMode() string    { return d.sslmode }

func (c *config) Db() IDbConfig {
	return c.db
}

type jwt struct {
	adminKey         string
	secretKey        string
	apikey           string
	accessExpiresAt  int
	refreshExpiresAt int
}

type IJwtConfig interface {
	SecretKey() []byte
	AdminKey() []byte
	Apikey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

func (j *jwt) SecretKey() []byte          { return []byte(j.secretKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) Apikey() []byte             { return []byte(j.apikey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}
