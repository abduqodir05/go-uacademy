package models

type Post struct{
	Id string `json:"id"`
	Title string `json:"title"`
	UserId int `json:"userId"`
	Description string `json:"descriptioner"`
}
type CreatePost struct {
	Id string `json:"id"`
	Title string `json:"title"`
	UserId int `json:"userId"`
	Description string `json:"descriptioner"`
}
type getAllPost struct {
	Id string `json:"id"`
	Title string `json:"title"`
	UserId int `json:"userId"`
	Description string `json:"descriptioner"`
}
type getByIdPost struct {
	Id string `json:"id"`
}
type getUserPost struct {
	Id string `json:"id"`
}