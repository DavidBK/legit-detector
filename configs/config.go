package configs

type Config struct {
	Port     int
	LogLevel string
}

const (
	DefaultPort     = 8080
	DefaultLogLevel = "info"
)

func NewConfig() *Config {
	return &Config{
		Port:     DefaultPort,
		LogLevel: DefaultLogLevel,
	}
}
