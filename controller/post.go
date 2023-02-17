package controller

import "go-uacademy/models"

func (c *Controller) CreatePost(req models.CreatePost) (id int, err error){
	
	id, err = c.store.Post().CreatePost(&req) 
	if err !=nil {
		return 0,err
	}
	return id, nil
}
func (c *Controller) GetByIdPost(id string) (Fpost *models.Post, err error){
	Fpost, err = c.store.Post().GetByIdPost(id)
	if err !=nil {
		return Fpost,err
	}
	return Fpost, nil
	
}

