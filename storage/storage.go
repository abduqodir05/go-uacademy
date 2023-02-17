package storage

import "go-uacademy/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Post() PostReopI
}
type UserRepoI interface {
	GetUserById(req *models.User) (models.User, error)
	Create(req *models.CreateUser) (id int, err error)
	Update(req *models.UpdateUser) (res models.User, err error)
	Delete(req *models.User) (res models.User, err error)
	GetList(req *models.GetListRequest) (*models.GetListResponse, error)
}
type PostReopI interface {
	CreatePost(req *models.CreatePost) (id int, err error)
	GetByIdPost(id string) (*models.Post, error)
}
