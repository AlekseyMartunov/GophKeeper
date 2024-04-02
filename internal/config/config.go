package config

type Config struct {
	secretKey            string
	postgresDSN          string
	salt                 string
	runAddr              string
	minioAccessKeyID     string
	minioSecretAccessKey string
	minioEndpoint        string
}

func NewConfig() *Config {
	return &Config{
		secretKey:            "thisis32bitlongpassphraseimusing",
		postgresDSN:          "postgres://admin:1234@localhost:5432/test?sslmode=disable",
		runAddr:              "127.0.0.1:8080",
		minioEndpoint:        "127.0.0.1:9090",
		minioAccessKeyID:     "minioServer",
		minioSecretAccessKey: "minioServer123",
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
