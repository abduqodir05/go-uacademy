package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-uacademy/models"
	"io/ioutil"
	"os"
)

type postRepo struct {
	fileName string
	file     os.File
}

func NewPostRepo(fileName string, file *os.File) *postRepo {
	return &postRepo{
		fileName: fileName,
		file:     *file,
	}
}

func (u *postRepo) CreatePost(req *models.CreatePost) (id int, err error) {
	var posts []*models.Post
	data, err := ioutil.ReadFile("./data/post.json")
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return 0, err
	}
	if len(posts) > 0 {
		posts = append(posts, &models.Post{
			Id:          req.Id,
			Title:       req.Title,
			UserId:      req.UserId,
			Description: req.Description,
		})
	} else {
		id = 1
		posts = append(posts, &models.Post{
			Id:          req.Id,
			Title:       req.Title,
			UserId:      req.UserId,
			Description: req.Description,
		})
	}

	body, err := json.MarshalIndent(posts, "", "   ")

	if err != nil {
		fmt.Println("adasda")

		return 0, nil
	}

	err = ioutil.WriteFile("./data/post.json", body, os.ModePerm)
	if err != nil {
		fmt.Println("adasda")
		return 0, err
	}

	return id, nil
}

func (u *postRepo) GetByIdPost(id string) (*models.Post, error) {
	var posts []*models.Post
	data, err := ioutil.ReadFile("./data/post.json")
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

	return &models.Post{}, errors.New("Post Not Found")

}
