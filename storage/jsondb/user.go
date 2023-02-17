package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-uacademy/models"
	"io/ioutil"
	"os"
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

func readFile() ([]models.User, error) {
	users := []models.User{}
	data, err := ioutil.ReadFile("./data/users.json")
	// err = json.NewDecoder(u.file).Decode(users)
	if err != nil {
		fmt.Println("Get Abdulqodir Update")

		return []models.User{}, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return []models.User{}, err
	}

	return users, nil

}

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

	users, err := readFile()
	if err != nil {
		return 0, err
	}
	if len(users) > 0 {
		users = append(users, models.User{
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
		users = append(users, models.User{
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

	body, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("asdsdf")
		return 0, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		fmt.Println("asdsdf")
		return 0, err
	}

	return id, nil
}

func (u *userRepo) Update(req *models.UpdateUser) (res models.User, err error) {
	users := []models.User{}
	// ============
	data, err := ioutil.ReadFile(u.fileName)
	// err = json.NewDecoder(u.file).Decode(users)
	if err != nil {
		fmt.Println("Get Abdulqodir Update")

		return models.User{}, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return models.User{}, err
	}

	for idx, v := range users {
		if users[idx].Id == v.Id {
			fmt.Println("SUCCESS: Updated the user")
			users[idx].First_name = v.First_name
			users[idx].Last_name = v.Last_name
			return
		}
	}
	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (u *userRepo) Delete(req *models.User) (res models.User, err error) {
	var users []models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		fmt.Println("Get list")

		return models.User{}, err
	}

	for ind, val := range users {
		if users[ind].Id == val.Id {
			fmt.Println("SUCCESS: Deleted the user")
			users = append(users[ind:], users[ind+1:]...)
			return
		}
	}
	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (u *userRepo) GetUserById(req *models.User) (models.User, error) {

	var users []models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		fmt.Println("Get list")

		return models.User{}, err
	}

	for ind, v := range users {
		if users[ind].Id == v.Id {
			return v, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func (u *userRepo) Read(req *models.CreateUser) (id int, err error) {

	var users []models.User
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
		fmt.Println("Get list")

		return nil, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Get list")
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
