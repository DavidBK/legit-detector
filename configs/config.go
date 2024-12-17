package configs

type Config struct {
	Port int
}

const (
	DefaultPort = 8080
)

func NewConfig() *Config {
	return &Config{
		Port: DefaultPort,
	}
}
