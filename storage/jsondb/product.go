package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-uacademy/models"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

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
func (u *ProductRepo) CreateProduct(req *models.CreateProduct) (id string, err error) {

	id = uuid.New().String()

	var product []*models.Product
	data, err := ioutil.ReadFile("./data/product.json")
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, &product)
	if err != nil {
		return "", err
	}
	product = append(product, &models.Product{
		Id:    id,
		Name:  req.Name,
		Price: req.Price,
	})

	body, err := json.MarshalIndent(product, "", "   ")

	if err != nil {
		fmt.Println("adasda")

		return "", nil
	}

	err = ioutil.WriteFile("./data/product.json", body, os.ModePerm)
	if err != nil {
		fmt.Println("adasda")
		return "", err
	}
	return id, nil
}
func (u *ProductRepo) UpdateProduct(req *models.UpdateProduct) (res *models.Product, err error) {

	id := uuid.New().String()

	var product []*models.Product
	data, err := ioutil.ReadFile("./data/product.json")
	if err != nil {
		fmt.Println("couldn't read file")
		return nil, err
	}
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("couldn't be unmarsharal")
		return nil, err
	}

	for ind, val := range product {
		if product[ind].Id == val.Id {
			product[ind].Id = id
			product[ind].Name = val.Name
			product[ind].Price = val.Price
		}
	}
	body, err := json.MarshalIndent(product, "", "   ")

	if err != nil {
		fmt.Println("adasda")

		return nil, err
	}

	err = ioutil.WriteFile("./data/product.json", body, os.ModePerm)
	if err != nil {
		fmt.Println("adasda")
		return nil, err
	}
	return res, nil
}

func (u *ProductRepo) DeleteProduct(req *models.DeleteProduct) (res *models.Product, err error) {

	var product []*models.Product
	data, err := ioutil.ReadFile("./data/product.json")
	if err != nil {
		fmt.Println("couldn't read file")
		return nil, err
	}
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("couldn't be unmarsharal")
		return nil, err
	}

	for ind, v := range product {
		if product[ind].Id == v.Id {
			product = append(product[ind:], product[ind+1:]...)
			return
		}
	}
	body, err := json.MarshalIndent(product, "", "   ")

	if err != nil {
		fmt.Println("couldn't be marshal")

		return nil, err
	}

	err = ioutil.WriteFile("./data/product.json", body, os.ModePerm)
	if err != nil {
		fmt.Println("couldn't read file")
		return nil, err
	}
	return res, nil
}

func (u *ProductRepo) GetByIdProduct(id string) (*models.Product, error) {
	var posts []*models.Product
	data, err := ioutil.ReadFile("./data/product.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, err
	}
	
	for _, val := range posts {
		if id == val.Id {
			return val, nil
		}
	}

	return &models.Product{}, errors.New("post Not Found")
}
