package config

type AppConfig struct {
	Name string `json:"name"`
}

var Config = &AppConfig{}

func init() {
	Config = &AppConfig{
		Name: "init config",
	}
}

func Load() *AppConfig {
	Config = &AppConfig{
		Name: "load config",
	}
	return Config
}
