package controller

import (
	"go-uacademy/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (c *Controller) UpdateUser(req *models.UpdateUser) (res *models.User, err error) {
	user, err := c.store.User().Update(res)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (c *Controller) DeleteUser(req *models.DeleteUser) (res *models.User, err error) {
	User, err := c.store.User().Delete(res)	
	if err != nil {
		return nil, err
	}

	return User, nil
}

// func (c *Controller) GetUserById(req *models.UserPrimaryKey) (*models.User, error) {

// 	users, err := c.store.User().Getlist(req)
// 	if err != nil {
// 		return &models.User{}, err
// 	}
// 	return users, nil
// }

// func (c *Controller) Getlist(req *models.GetListRequest) (*models.GetListResponse, error) {

// 	user, err := c.store.User().Getlist(req)
// 	if err != nil {
// 		return &models.GetListResponse{}, err
// 	}

// 	return user, nil
// }
