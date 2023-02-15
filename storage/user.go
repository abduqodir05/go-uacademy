package storage

import (
	"encoding/json"
	"go-uacademy/models"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(f string, file *os.File) *userRepo {
	return &userRepo{
		fileName: f,
		file:     file,
	}
}

func (u *userRepo) Getlist(search string) ([]models.User, error) {
	var users []models.User

	data, err := ioutil.ReadFile("users.json")
	if err != nil {

		return users, err
	}
	err = json.Unmarshal(data, &users)

	if err != nil {
		log.Println("invalid user")
		return nil, err
	}

	for _, v := range users {
		if strings.Contains(v.First_name, search) || strings.Contains(v.Last_name, search) {
			users = append(users, v)
		}
	}
	return users, nil
}

func (u *userRepo) Create(req *models.User) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}
	if len(users) > 0 {
		users = append(users, &models.User{
			Id:          req.Id,
			First_name:  req.First_name,
			Last_name:   req.Last_name,
			Gender:      req.Gender,
			Address:     req.Address,
			Card_number: req.Card_number,
			Birthday:    req.Birthday,
			Profession:  req.Profession,
		})
	} else {
		id = 1
		users = append(users, &models.User{
			Id:          req.Id,
			First_name:  req.First_name,
			Last_name:   req.Last_name,
			Gender:      req.Gender,
			Address:     req.Address,
			Card_number: req.Card_number,
			Birthday:    req.Birthday,
			Profession:  req.Profession,
		})
	}

	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile("/data/users.json", body, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepo) Update(req *models.User) error{
	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return err
	}
	return nil

	for idx, v := range users {
		if users[idx].Id == v.Id {
			users[idx].First_name = v.First_name
			users[idx].Last_name = v.Last_name
		}
	}
	body, err := json.MarshalIndent(req, "", "   ")

	err = ioutil.WriteFile("users.json", body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}


func (u *userRepo) Read(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	return id, nil
}
