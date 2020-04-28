package store

//
type Config struct {
	DatabaseURL  string
	DataBaseName string
}

//
func NewConfig(url string, name string) *Config {
	return &Config{
		DatabaseURL:  url,
		DataBaseName: name,
	}
}
