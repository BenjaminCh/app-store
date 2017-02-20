package interfaces

// IConfiguration : Configuration interface.
// Not to rely on a concrete configuration implementation.
type IConfiguration interface {
	GetConfig(key string, defaultValue interface{}) interface{}
	GetConfigString(key string, defaultValue string) string
	GetConfigBool(key string, defaultValue bool) bool
	GetConfigUInt(key string, defaultValue uint) uint
}

// ConfigurationManager : contains the configuration interactor.
type ConfigurationManager struct {
	ConfigurationInteractor IConfiguration
}

// GetConfig allows to retrieve a configuration value for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfig(key string, defaultValue interface{}) interface{} {
	return c.ConfigurationInteractor.GetConfig(key, defaultValue)
}

// GetConfigString allows to retrieve a configuration String for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigString(key string, defaultValue string) string {
	return c.ConfigurationInteractor.GetConfig(key, defaultValue).(string)
}

// GetConfigBool allows to retrieve a configuration Boolean for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigBool(key string, defaultValue bool) bool {
	return c.ConfigurationInteractor.GetConfig(key, defaultValue).(bool)
}

// GetConfigUInt allows to retrieve a configuration UInt for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigUInt(key string, defaultValue uint) uint {
	return c.ConfigurationInteractor.GetConfig(key, defaultValue).(uint)
}
