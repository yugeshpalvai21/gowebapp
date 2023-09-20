package config

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// return config instance
func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "yugeshpalvai",
		Password: "password",
		Name:     "gowebapp",
	}
}
