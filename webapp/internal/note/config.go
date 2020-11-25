package note

// Config tarantool
type Config struct {
	TarantoolAddr string `toml:"tarantool_addr"`
	User          string `toml:"user"`
	Pass          string `toml:"pass"`
}

// NewConfig tarantool
func NewConfig() *Config {
	return &Config{}
}
