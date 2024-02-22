package config

type Config struct {
	secretKey string
}

func NewConfig() *Config {
	return &Config{secretKey: "thisis32bitlongpassphraseimusing"}
}

func (c *Config) SecretKey() string {
	return c.secretKey
}
