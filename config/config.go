package config

type Config struct {
	Path string
	UserFileName string
	PostFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cd
	cfg.UserFileName = "/users.json"

	return cfg
}
