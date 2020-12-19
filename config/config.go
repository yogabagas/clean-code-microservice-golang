package config

import (
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Source string

const (
	DefaultType     = "json"
	DefaultFileName = "config"

	EnvConsulHostKey = "GOCOF_CONSUL"
	EnvTypeKey       = "GOCOF_TYPE"
	EnvFileNameKey   = "GOCOF_FILENAME"
	EnvPrefixKey     = "GOCONF_ENV_PREFIX"

	SourceEnv    Source = "env"
	SourceConsul Source = "consul"
)

var (
	typ    = DefaultType
	fname  = DefaultFileName
	prefix string

	c    *viper.Viper
	dirs = []string{
		".",
		"/app/config",
	}

	errEnv, errFile, errConsule error
)

func Configure() *viper.Viper {
	// Load file .env
	if err := godotenv.Load(); err != nil {
		errEnv = errors.Cause(err)
	}

	if v := os.Getenv(EnvTypeKey); v != "" {
		typ = v
	}

	if v := os.Getenv(EnvFileNameKey); v != "" {
		fname = v
	}

	if v := os.Getenv(EnvPrefixKey); v != "" {
		prefix = v
	}

	c = viper.New()
	c.SetConfigType(typ)
	c.SetConfigFile(fname)
	if len(prefix) > 0 {
		c.SetEnvPrefix(prefix)
	}
	c.AutomaticEnv()

	if ch := os.Getenv(EnvConsulHostKey); ch != "" {
		if err := c.AddRemoteProvider("consul", ch, fname); err != nil {
			errConsule = errors.Cause(err)
		} else {
			// create var func to modify backoff operation
			connect := func() error { return c.ReadRemoteConfig() }
			// create var func to modify backoff notify
			notify := func(err error, t time.Duration) { log.Println("[goconf", err.Error(), t) }
			// create new instance of backoff
			b := backoff.NewExponentialBackOff()
			b.MaxElapsedTime = (2 * time.Minute)

			if err := backoff.RetryNotify(connect, b, notify); err != nil {
				log.Printf("[goconf] giving up connecting to remote config ")
				errConsule = errors.Cause(err)
			}
		}
	} else {
		errConsule = errors.New("failed loading remote source; ENV not defined")
	}

	// last try with file dirs

	for _, d := range dirs {
		c.AddConfigPath(d)
	}

	if err := c.ReadInConfig(); err != nil {
		errFile = errors.Cause(err)
	}

	return c
}

func GetStringSlice(s string) []string {
	return c.GetStringSlice(s)
}
