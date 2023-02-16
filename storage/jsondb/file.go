package jsondb

import (
	"go-uacademy/config"
	"go-uacademy/storage"
	"os"
)

type Store struct {
	product *ProductRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {

	ProductFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		product: NewProductRepo(cfg.Path+cfg.ProductFileName, ProductFile),
	}, nil

}

func (s *Store) CloseDB() {
	s.product.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.product
}
