package interfaces

type IConfiguration interface {
	GetConfig(key string) interface{}
	GetConfigString(key string) string
	GetConfigBool(key string) bool
	GetConfigUInt(key string) uint
}

type ConfigurationManager struct {
	ConfigurationInteractor IConfiguration
}

// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigString(key string) string {
	return c.ConfigurationInteractor.GetConfigString(key)
}

// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigBool(key string) bool {
	return c.ConfigurationInteractor.GetConfigBool(key)
}

// Implements function from IConfiguration interface
func (c ConfigurationManager) GetConfigUInt(key string) uint {
	return c.ConfigurationInteractor.GetConfigUInt(key)
}
