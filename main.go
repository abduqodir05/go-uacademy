package main

import (
	"fmt"
	"go-uacademy/config"
	"go-uacademy/controller"
	"go-uacademy/models"
	"go-uacademy/storage"
	"log"
)

func main() {

	cfg := config.Load()
	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	c := controller.NewController(&cfg, store)

	// create user
	id, err := c.CreateUser(
		&models.CreateUser{
			Id:          "1",
			First_name:  "Abduqodir",
			Last_name:   "Musayev",
			Gender:      "male",
			Card_number: "7289057348573",
			Birthday:    "2005-10-23",
			Profession:  "student",
		},
	)

	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(id)

	// update user
	user, err := c.UpdateUser(
		&models.UpdateUser{
			Id:          "c57aa672-902f-44c8-af9d-dfa02f62541a",
			First_name:  "Abduqodir",
			Last_name:   "Musayev",
			Gender:      "male",
			Card_number: "321398042734275094",
			Birthday:    "2005-10-23",
			Profession:  "Team lead at Amazon",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	
	// delete user
	deleted_user, err := c.DeleteUser(
		&models.DeleteUser{
			First_name:  "Abduqodir",
			Last_name:   "Musayev",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deleted_user)
	
	// get_by_id

	user, err = c.GetUserById(
		&models.UserPrimaryKey{
			Id: 1,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GetUserById:", user)

}
