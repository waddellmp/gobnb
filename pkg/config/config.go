package config

// Configuration
type appConfig struct {
	Port     string
	UseCache bool
}

var AppConfig appConfig = appConfig{
	Port:     ":8080",
	UseCache: false,
}
