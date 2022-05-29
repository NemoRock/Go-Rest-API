package apiserver

type Config struct {
	BimAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BimAddr:  ":8080",
		LogLevel: "debug",
	}
}
