package storage

import "go-uacademy/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (int, error)
	GetPkey(*models.UserPrimaryKey) (*models.User, error)
	Update(models.UpdateUser) (int error)
	// Getlist(*models.Getlist) (string error)
}
