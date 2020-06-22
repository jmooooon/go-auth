package auth

// Config ...
type Config struct {
	Port string `toml:"port"`
	DB   string `toml:"dburl"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config {
		Port: ":8080",
		DB: "debug",
	}
}
