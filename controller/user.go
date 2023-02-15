package controller

import (
	"go-uacademy/models"
	"go-uacademy/storage"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	// id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (c *Controller) UpdateUser(req *models.CreateUser) (id int, err error) {

	// id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Controller) Getlist(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User.Read(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}