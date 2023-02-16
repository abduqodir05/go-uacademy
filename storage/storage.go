package storage

import "go-uacademy/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}
type UserRepoI interface {
	GetUserById(req *models.User) (*models.User,  error)
	Create(req *models.CreateUser) (id int, err error)
	Update(req *models.User) (res *models.User, err error)
	Delete(req *models.User) (res *models.User, err error)
	GetList(req *models.GetListRequest) (*models.GetListResponse, error)
}