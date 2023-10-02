package config

import "os"

type Config struct {
	MongoUri string
	DBName   string
}

func LoadConfig() *Config {
	return &Config{
		MongoUri: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
