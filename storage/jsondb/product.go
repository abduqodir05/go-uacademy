package jsondb

import "os"

type ProductRepo struct {
	fileName string
	file     *os.File
}
type UserRepo struct {
	fileName string
	file     *os.File
}

func NewProductRepo(fileName string, file *os.File) *ProductRepo {
	return &ProductRepo{
		fileName: fileName,
		file:     file,
	}
}
func NewUserRepo(fileName string, file *os.File) *UserRepo {
	return &UserRepo{
		fileName: fileName,
		file:     file,
	}
}
