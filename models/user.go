package models

type User struct {
	Id          string   `json:"id"`
	First_name  string   `json:"first_name"`
	Last_name   string   `json:"last_name"`
	Gender      string   `json:"gender"`
	Card_number string   `json:"card_number"`
	Birthday    string   `json:"birthday"`
	Profession  string   `json:"profession"`
	Address     Address  `json:"address"`
	Friends     []Friend `json:"friends"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type Friend struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
}
type UpdateUser struct {
	Id          string   `json:"id"`
	First_name  string   `json:"first_name"`
	Last_name   string   `json:"last_name"`
	Gender      string   `json:"gender"`
	Card_number string   `json:"card_number"`
	Birthday    string   `json:"birthday"`
	Profession  string   `json:"profession"`
	Address     Address  `json:"address"`
	Friends     []Friend `json:"friends"`
}
type CreateUser struct {
	Id          string   `json:"id"`
	First_name  string   `json:"first_name"`
	Last_name   string   `json:"last_name"`
	Gender      string   `json:"gender"`
	Card_number string   `json:"card_number"`
	Birthday    string   `json:"birthday"`
	Profession  string   `json:"profession"`
	Address     Address  `json:"address"`
	Friends     []Friend `json:"friends"`
}
type UserPrimaryKey struct {
	Id int `json:"id"`
}
type DeleteUser struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
type GetListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
type GetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}