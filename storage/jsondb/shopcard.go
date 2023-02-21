package jsondb

import "os"

type ShopCardRepo struct {
	fileName string
	file     *os.File
}

func NewShopCardRepo(fileName string, file *os.File) *ShopCardRepo {
	return &ShopCardRepo{
		fileName: fileName,
		file:     file,
	}
}

