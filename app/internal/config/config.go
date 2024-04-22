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
	migrationPath        string `env:"MIGRATION_PATH"`
}

func NewConfig() *Config {
	return &Config{
		secretKey:            "default_secret_key",
		salt:                 "some_salt",
		postgresDSN:          "postgres://admin:1234@localhost:5432/test?sslmode=disable",
		runAddr:              "127.0.0.1:8080",
		minioEndpoint:        "127.0.0.1:9001",
		minioAccessKeyID:     "minioServer",
		minioSecretAccessKey: "minioServer123",
		migrationPath:        "./migrations",
	}
}

func (c *Config) ParseFlags() {
	if key, ok := os.LookupEnv("SECRET_KEY"); ok {
		c.secretKey = key
	}

	if key, ok := os.LookupEnv("POSTGRES_DSN"); ok {
		c.postgresDSN = key
	}

	if key, ok := os.LookupEnv("SALT"); ok {
		c.salt = key
	}

	if key, ok := os.LookupEnv("RUN_ADDRESS"); ok {
		c.runAddr = key
	}

	if key, ok := os.LookupEnv("MINIO_ACCESS_KEY_ID"); ok {
		c.minioSecretAccessKey = key
	}

	if key, ok := os.LookupEnv("MINIO_SECRET_KEY_ID"); ok {
		c.minioSecretAccessKey = key
	}

	if key, ok := os.LookupEnv("MINIO_ENDPOINT"); ok {
		c.minioEndpoint = key
	}

	if key, ok := os.LookupEnv("MIGRATION_PATH"); ok {
		c.migrationPath = key
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

func (c *Config) MigrationPath() string {
	return c.migrationPath
}
