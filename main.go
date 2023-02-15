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

type User struct {
	Id          string   `json:"id"`
	First_name  string   `json:"first_name"`
	Last_name   string   `json:"last_name"`
	Gender      string   `json:"gender"`
	Card_number string   `json:"card_number"`
	Birthday    string   `json:"birthday"`
	Profession  string   `json:"profession"`
	Address     Address  `json:"address"`
	Friends     []friend `json:"friends"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type friend struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
}

func Update(req []User, u User) {
	for idx, v := range req {
		if v.Id == u.Id {
			req[idx].First_name = u.First_name
			req[idx].Last_name = u.Last_name
			req[idx].Gender = u.Gender
			req[idx].Card_number = u.Card_number
			req[idx].Birthday = u.Birthday
			req[idx].Profession = u.Profession
			req[idx].Address = u.Address
			req[idx].Friends = u.Friends
		}
	}
	body, err := json.MarshalIndent(req, "", "   ")

	err = ioutil.WriteFile("users.json", body, os.ModePerm)
	if err != nil {
		return
	}

}
func Delete(req []User, id string) {

	for ind, v := range req {
		if v.Id == id {
			req = append(req[:ind], req[ind+1:]...)
		}
	}
	body, err := json.MarshalIndent(req, "", "   ")

	err = ioutil.WriteFile("users.json", body, os.ModePerm)
	if err != nil {
		return
	}
}

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
	
	
	
	b := controller.NewController(&cfg, store)

	id, err =  b.UpdateUser(
		&models.Update{
			
		}
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
