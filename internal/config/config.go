package config

type Config struct {
	secretKey   string
	postgresDSN string
	salt        string
	runAddr     string
}

func NewConfig() *Config {
	return &Config{
		secretKey:   "thisis32bitlongpassphraseimusing",
		postgresDSN: "postgres://admin:1234@localhost:5432/test?sslmode=disable",
		runAddr:     "127.0.0.1:8080",
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
