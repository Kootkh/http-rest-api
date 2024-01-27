package apiserver

// Config - конфигурация APIServer
type Config struct {

	// Адрес, на котором запущен APIServer
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig - конфигурация APIServer
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8084",
		LogLevel: "debug",
	}
}
