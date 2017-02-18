package infrastructure

import (
	"log"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// ConfigHandler dealing with configuration keys.
type ConfigHandler struct {
	config *viper.Viper
}

// NewConfig allows to create a new configuration handler based on viper.
func NewConfig() *ConfigHandler {

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		viper.AddConfigPath(path.Dir(path.Dir(filename)))
	}
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("initConfig : No configuration file found, using the default config.")
	}

	// Env variables
	viper.SetEnvPrefix("ALGOLIA_APP_STORE")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Common variables
	viper.SetDefault("version", "0.0.1.0")
	viper.SetDefault("server.port", "8888")
	viper.SetDefault("application.debug", false)

	// Algolia specific part
	viper.SetDefault("algolia.applicationID", "NOT_SET")
	viper.SetDefault("algolia.apiKey", "NOT_SET")

	return &ConfigHandler{config: viper.GetViper()}
}

func (c ConfigHandler) GetConfig(key string) interface{} {
	return c.config.Get(key)
}

// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigString(key string) string {
	return c.config.Get(key).(string)
}

// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigBool(key string) bool {
	return c.config.Get(key).(bool)
}

// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigUInt(key string) uint {
	return c.config.Get(key).(uint)
}
