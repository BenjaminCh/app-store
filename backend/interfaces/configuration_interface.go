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
