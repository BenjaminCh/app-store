package interfaces

// IConfiguration : Configuration interface.
// Not to rely on a concrete configuration implementation.
type IConfiguration interface {
	GetConfig(key string) interface{}
	GetConfigString(key string) string
	GetConfigBool(key string) bool
	GetConfigUInt(key string) uint
}

// ConfigurationManager : contains the configuration interactor.
type ConfigurationManager struct {
	ConfigurationInteractor IConfiguration
}

// GetConfig allows to retrieve a configuration value for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfig(key string) interface{} {
	return c.ConfigurationInteractor.GetConfig(key)
}

// GetConfigString allows to retrieve a configuration String for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigString(key string) string {
	return c.ConfigurationInteractor.GetConfig(key).(string)
}

// GetConfigBool allows to retrieve a configuration Boolean for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigBool(key string) bool {
	return c.ConfigurationInteractor.GetConfig(key).(bool)
}

// GetConfigUInt allows to retrieve a configuration UInt for a given key.
// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigUInt(key string) uint {
	return c.ConfigurationInteractor.GetConfig(key).(uint)
}
