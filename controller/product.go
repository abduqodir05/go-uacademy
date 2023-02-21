package controller

import (
	"go-uacademy/models"
)

func (c *Controller) CreateProduct(req *models.CreateProduct) (id string, err error) {

	id, err = c.store.Product().CreateProduct(req)
	if err != nil {
		return "", err
	}

	return id, nil
}
func (c *Controller) UpdateProduct(req *models.UpdateProduct) (res *models.Product, err error) {
	product, err := c.store.Product().UpdateProduct(req)
	if err != nil {
		return &models.Product{}, err
	}

	return product, nil
}

func (c *Controller) DeleteProduct(req *models.DeleteProduct) (res *models.Product, err error) {
	product, err := c.store.Product().DeleteProduct(req)
	if err != nil {
		return &models.Product{}, err
	}

	return product, nil
}
func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (prduct *models.Product, err error){
	
	prduct, err = c.store.Product().GetByIdProduct(req.Id)
	if err != nil  {
		return prduct, err
	}
	return prduct, nil

}