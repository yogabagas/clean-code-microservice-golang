package cmd

import (
	"database/sql"
	"fmt"
	"my-github/clean-code-microservice-golang/config"
	"my-github/clean-code-microservice-golang/infrastructure/datastore"

	log "github.com/sirupsen/logrus"

	"github.com/fsnotify/fsnotify"

	"github.com/evalphobia/logrus_sentry"

	redis "github.com/go-redis/redis/v7"
)

var (
	CfgMySqlUser                  = "database.mysql.dbuser"
	CfgMysqlPwd                   = "database.mysql.dbpwd"
	CfgMysqlPort                  = "database.mysql.port"
	CfgMysqlDBName                = "database.mysql.dbname"
	CfgRedisHost                  = "database.redis.host"
	CfgRedisPass                  = "database.redis.password"
	CfgRedisDB                    = "database.redis.db"
	CfgKafkaGroup                 = "kafka.group_id"
	CfgKafkaHost                  = "kafka.host"
	CfgKafkaProtocol              = "kafka.security_protocol"
	CfgKafkaMechanisms            = "kafka.sasl_mechanisms"
	CfgKafkaKey                   = "kafka.sasl_username"
	CfgKafkaSecret                = "kafka.sasl_password"
	CfgKafkaTopic                 = "kafka.topics"
	CfgNewRelicKey                = "newrelic.key"
	CfgNewRelicDebug              = "newrelic.debug"
	CfgMongoURI                   = "database.mongo.uri"
	CfgMongoDB                    = "database.mongo.db"
	CfgSentryKey                  = "sentry.key"
	TelemetryID                   = "newrelic.id"
	CfgNewrelicSlowQueryThreshold = "newrelic.slowquery.threshold"
	CfgNewrelicSlowQueryEnabled   = "newrelic.slowquery.enabled"
	CfgDkronHost                  = "dkron.host"
)

var (
	db     *sql.DB
	rdb    *redis.Ring
	logger *log.Logger
)

func init() {
	c := config.Configure()

	go func() {
		c.WatchConfig()
		c.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("config file changed %v", e.Name)
		})
	}()

	db, _ = InitDB()
	rdb = InitRedis()

	logger = InitLogger()
}

func InitLogger() *log.Logger {
	log.SetFormatter(&log.JSONFormatter{})
	l := log.StandardLogger()
	if dsn := config.GetString(CfgSentryKey); len(dsn) > 0 {
		hook, err := logrus_sentry.NewSentryHook(dsn, []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		})
		if err == nil {
			hook.StacktraceConfiguration.Enable = true
			l.Hooks.Add(hook)
		}
	}
	return nil
}

func InitDB() (db *sql.DB, err error) {
	conn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", config.GetString(CfgMySqlUser), config.GetString(CfgMysqlPwd), config.GetInt(CfgMysqlPort), config.GetString(CfgMysqlDBName))
	return datastore.NewDB("mysql", conn)
}

func InitRedis() (rdb *redis.Ring) {
	return datastore.NewRedisClient(config.GetStringMapString(CfgRedisHost), config.GetString(CfgRedisPass), config.GetInt(CfgRedisDB))
}
