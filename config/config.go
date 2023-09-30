package config

import "os"

type Config struct {
	MongoUri string
}

func LoadConfig() *Config {
	return &Config{
		MongoUri: os.Getenv("MONGO_URI"),
	}
}
