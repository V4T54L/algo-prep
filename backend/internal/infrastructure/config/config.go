package config

type Config struct {
	Port         string
	Env          string
	DatabaseURL  string
	JWTSecret    string
	SessionKey   string
	LogLevel     string
	TemplatePath string
}

func Load() (*Config, error) {
	// Implementation would load from env files
	return &Config{
		Port:         ":8080",
		Env:          "development",
		DatabaseURL:  "postgres://user:pass@localhost/db",
		JWTSecret:    "your-secret-key",
		SessionKey:   "session-secret-key",
		LogLevel:     "debug",
		TemplatePath: "./web/templates",
	}, nil
}
