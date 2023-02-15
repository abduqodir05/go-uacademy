package controller

import (
	"go-uacademy/models"
	"go-uacademy/storage"
)

func (c *Controller) CreateUser(*models.CreateUser) (id int, err error) {

	// id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (c *Controller) UpdateUser(req *models.UpdateUser) (res *models.User, err error) {
	user, err := c.store.User().Update(req)
	if err != nil {
	  return nil, err
	}
  
	return user, nil
  }

func (c *Controller) Getlist(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User().Read(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}