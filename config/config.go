package config

var (
	logger *Logger
)

func InitConfig() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
