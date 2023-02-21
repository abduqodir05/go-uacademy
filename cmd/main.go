package main

import (
	"fmt"
	"go-uacademy/config"
	"go-uacademy/controller"
	"go-uacademy/models"
	"go-uacademy/storage/jsondb"
	"log"
	// "github.com/google/uuid"
)

func main() {
	// asd = uuid.New().String()
	
	cfg := config.Load()
	store, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	
	c := controller.NewController(&cfg, store)
	
	id, err := c.CreateProduct(
		&models.CreateProduct{
			Id:    "2",
			Name:  "grapes",
			Price: 41.9,
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(id)

	productUpdate, err := c.UpdateProduct(
		&models.UpdateProduct{
			Id: "121",
			Name: "grapes",
			Price: 17.8,
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(productUpdate)

	productDelete, err := c.DeleteProduct(
		&models.DeleteProduct{
			Id: 153,
			Name: "orange",
			Price: 14.9,
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(productDelete)

	GetByIdProduct, err := c.GetByIdProduct(
		&models.ProductPrimaryKey{
			Id: "1",
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(GetByIdProduct)
}
