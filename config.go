package apollo

import "time"

type Config struct {
	Url   string
	Token string

	Timeout time.Duration
	Logger  Logger
}

func DefaultConfig() Config {
	return Config{
		Timeout: 1 * time.Minute,
		Logger:  DefaultLogger,
	}
}
