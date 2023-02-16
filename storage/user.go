package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-uacademy/models"
	"io/ioutil"
	// "log"
	"os"
	// "strings"
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

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

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

func (u *userRepo) Update(req *models.User) (res *models.User, err error) {
	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}

	for idx, v := range users {
		if users[idx].Id == v.Id {
			fmt.Println("SUCCESS: Updated the user")
			users[idx].First_name = v.First_name
			users[idx].Last_name = v.Last_name
		}
	}
	body, err := json.MarshalIndent(req, "", "   ")

	err = ioutil.WriteFile("/data/users.json", body, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userRepo) Delete(req *models.User) (res *models.User, err error) {
	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}

	for ind, val := range users {
		if users[ind].Id == val.Id {
			fmt.Println("SUCCESS: Deleted the user")
			users = append(users[ind:], users[ind+1:]...)
		}
	}
	body, err := json.MarshalIndent(req, "", "   ")
	
	err = ioutil.WriteFile("/data/users.json", body, os.ModePerm)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (u *userRepo) GetUserById(req *models.User) (*models.User,  error) {
	
	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}
	
	for ind, v := range users {
		if users[ind].Id == v.Id {
			return v, nil
		}
	}
	
	body, err := json.MarshalIndent(req, "", "   ")

	err = ioutil.WriteFile("/data/users.json", body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &models.User{}, errors.New("user not found")
}

func (u *userRepo) Read(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	return id, nil
}


func (u *userRepo) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {
	users := make([]models.User, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	if req.Limit+req.Offset > len(users) {
		if req.Offset > len(users) {
			return &models.GetListResponse{
				Count: len(users),
				Users: []models.User{},
			}, nil
		}

		return &models.GetListResponse{
			Count: len(users),
			Users: users[req.Offset:],
		}, nil
	}

	response := &models.GetListResponse{
		Count: len(users),
		Users: users[req.Offset : req.Limit+req.Offset],
	}

	return response, nil
}