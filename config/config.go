package config

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func GetConfig() Config {
	return Config{
		DBHost:     "localhost",
		DBUser:     "postgres",
		DBPassword: "Kantapon4444#",
		DBName:     "taskmanager_db",
		DBPort:     "5432",
	}
}
