package apiserver

// Config ...
type Config struct {
	BindAddr  string `toml:"bind_addr"`
	StaticDir string `toml:"static_dir"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
