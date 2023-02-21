package config

type Config struct {
	Path string

	UserFileName string
	ProductFileName string
	ShopcardFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/users.json"
	cfg.ProductFileName = "/product.json"
	cfg.ShopcardFileName = "/product.json"

	return cfg
}

