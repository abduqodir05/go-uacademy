package config

type Config struct {
	Path string

	UserFileName string
	ProductFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/customer.json"
	cfg.UserFileName = "/product.json"

	return cfg
}

