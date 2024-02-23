package config

type Config struct {
	secretKey   string
	postgresDSN string
	salt        string
}

func NewConfig() *Config {
	return &Config{
		secretKey:   "thisis32bitlongpassphraseimusing",
		postgresDSN: "postgres://admin:1234@localhost:5432/test?sslmode=disable",
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
