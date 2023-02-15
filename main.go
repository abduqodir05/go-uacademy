package main

import (
	"encoding/json"
	"fmt"
	"go-uacademy/config"
	"go-uacademy/controller"
	"go-uacademy/models"
	"go-uacademy/storage"
	"io/ioutil"
	"log"
	"os"
)



func main() {
	// var users []User
	// data, err := ioutil.ReadFile("users.json")
	// if err != nil {
	// 	log.Println("invalid users")
	// 	return
	// }
	// // fmt.Println(string(data))
	// err = json.Unmarshal(data, &users)
	// if err != nil {
	// 	log.Println("invalid user")
	// 	return
	// }

	cfg := config.Load()
	store, _ := storage.NewFileJson(&cfg)
	res, err := store.User.Getlist("Yvette")
	if err != nil {
		log.Println("invalid User")
		return
	}
	log.Println(res)

	c := controller.NewController(&cfg, store)
	
	id, err := c.CreateUser(
		&models.CreateUser{
			Id: "1",
			First_name: "Abduqodir",
			Last_name: "Musayev",
			Gender: "male",
			Card_number: "7289057348573",
			Birthday: "2005-10-23",
			Profession: "student",
		},
	)
	
	if err != nil {
		log.Println("error while CreateUser:",err.Error())
		return
	}
	fmt.Println(id)
	
	
	
	_, err = c.UpdateUser(
		&models.UpdateUser{
		  First_name: "asd",
		  Last_name: "bbb",
		},
	  )
	
	
	// user := User{
	// 	Id: "c57aa672-902f-44c8-af9d-dfa02f62541a",
	// 	First_name: "Abduqodir",
	// 	Last_name: "Musayev",
	// 	Gender: "male",
	// 	Card_number: "321398042734275094",
	// 	Birthday: "2005-10-23",
	// 	Profession: "Team lead at Amazon",
	// }
	// log.Println(users)
	// id:="c57aa672-902f-44c8-af9d-dfa02f62541a"
	// Delete(users,id)
	// fmt.Printf("%+v",id)
}
