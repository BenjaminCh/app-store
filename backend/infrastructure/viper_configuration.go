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

// NewViperConfig allows to create a new configuration handler based on viper.
func NewViperConfig() *ConfigHandler {

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
	viper.SetEnvPrefix("APPSTORE")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Common variables
	viper.SetDefault("version", "0.0.1.0")
	viper.SetDefault("server.port", "5000")
	viper.SetDefault("application.debug", false)

	// Algolia specific part
	viper.SetDefault("algolia.applicationID", "NOT_SET")
	viper.SetDefault("algolia.apiKey", "NOT_SET")
	viper.BindEnv("algolia.applicationID") // Key set as environnement var on the server
	viper.BindEnv("algolia.apiKey")        // Key set as environnement var on the server
	// Indexes
	viper.SetDefault("algolia.indexes.apps", "apps")

	return &ConfigHandler{config: viper.GetViper()}
}

// GetConfig allows to retrieve a configuration value for a given key.
// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfig(key string, defaultValue interface{}) interface{} {
	if c.config.IsSet(key) == false {
		return defaultValue
	}
	return c.config.Get(key)
}

// GetConfigString allows to retrieve a configuration String for a given key.
// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigString(key string, defaultValue string) string {
	return c.GetConfig(key, defaultValue).(string)
}

// GetConfigBool allows to retrieve a configuration Boolean for a given key.
// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigBool(key string, defaultValue bool) bool {
	return c.GetConfig(key, defaultValue).(bool)
}

// GetConfigUInt allows to retrieve a configuration UInt for a given key.
// Implements function from IConfiguration interface
func (c ConfigHandler) GetConfigUInt(key string, defaultValue uint) uint {
	return c.GetConfig(key, defaultValue).(uint)
}
