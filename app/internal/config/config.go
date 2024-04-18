package config

import "os"

type Config struct {
	secretKey            string `env:"SECRET_KEY"`
	postgresDSN          string `env:"POSTGRES_DSN"`
	salt                 string `env:"SALT"`
	runAddr              string `env:"RUN_ADDRESS"`
	minioAccessKeyID     string `env:"MINIO_ACCESS_KEY_ID"`
	minioSecretAccessKey string `env:"MINIO_SECRET_KEY_ID"`
	minioEndpoint        string `env:"MINIO_ENDPOINT"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ParseFlags() {
	if key, ok := os.LookupEnv("SECRET_KEY"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("POSTGRES_DSN"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("SALT"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("RUN_ADDRESS"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("MINIO_ACCESS_KEY_ID"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("MINIO_SECRET_KEY_ID"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("MINIO_ENDPOINT"); ok {
		c.runAddr = key
	}

}

func (c *Config) SecretKey() string {
	return c.secretKey
}

func (c *Config) PostgresDSN() string {
	return c.postgresDSN
}

func (c *Config) Salt() string {
	return c.salt
}

func (c *Config) RunAddr() string {
	return c.runAddr
}

func (c *Config) MinioAccessKeyID() string {
	return c.minioAccessKeyID
}

func (c *Config) MinioSecretAccessKey() string {
	return c.minioSecretAccessKey
}

func (c *Config) MinioEndpoint() string {
	return c.minioEndpoint
}
