package storage

import "go-uacademy/models"

type StorageI interface {
	CloseDB()
	// User() UserRepoI
	Product() ProductRepoI
}

// type UserRepoI interface {
	// CreateUser(req *models.CreateUser) (id string, err error)
	// GetPkey(*models.UserPrimaryKey) (*models.User, error)
// }
type ProductRepoI interface {
	CreateProduct(req *models.CreateProduct) (id string, err error)
	UpdateProduct(req *models.UpdateProduct) (res *models.Product, err error)
	DeleteProduct(req *models.DeleteProduct) (res *models.Product, err error)
	// GetByIdProduct(req *models.ProductPrimaryKey) (id string, err error)

	GetByIdProduct(id string) (*models.Product, error)

}
