package config

type Config struct {
    fileStorageDir string
	listenerHost   string
	listenerPort   string
	dbHost         string
	dbPort         string
	database       string
	dbUser         string
	dbPassword     string
	testVersion    string
	authCookieName string
	authSecret     string
}

var *configInstance = nil

func (cfg Config) GetFileStorageName() string {
	return cfg.fileStorageDir
}

func (cfg Config) GetDBHost() string {
	return cfg.dbHost
}

func (cfg Config) GetDBPort() string {
	return cfg.dbPort
}

func (cfg *Config) GetDatabase() string {
	return cfg.database
}

func (cfg Config) GetDBUser() string {
	return cfg.dbUser
}

func (cfg Config) GetDBPassword() string {
	return cfg.dbPassword
}

func (cfg Config) IsTest() bool {
	return cfg.testVersion == "test"
}

func (cfg Config) GetAuthCookieName() string {
	return cfg.authCookieName
}

func (cfg Config) GetAuthSecret() string {
	return cfg.authSecret
}

func (cfg Config) GetListenerHost() string {
	return cfg.listenerHost
}

func (cfg Config) GetListenerPort() string {
	return cfg.listenerPort
}



func init() error {

}

func GetConfigInstance() (*Config, error) {
    if configInstance == nil {
        err := init() 
        if err != nil {
            return _, err
        }
    }
    return configInstance, nil
}