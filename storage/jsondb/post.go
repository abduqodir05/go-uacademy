package jsondb

import (
	"encoding/json"
	"errors"
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

func (u *userRepo) CreatePost(req *models.Post) (id int, err error) {
	var posts []*models.Post
	err = json.NewDecoder(u.file).Decode(&posts)
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

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepo) GetById(req *models.Post) (*models.Post, error){
	var posts []*models.Post
	err := json.NewDecoder(u.file).Decode(&posts)
	if err != nil {
		return nil, err
	}

	for ind, val := range posts {
		if posts[ind].Id == val.Id{
			return val, nil
		}
	}
	body, err := json.MarshalIndent(req, "" , "   ")

	err = ioutil.WriteFile(u.fileName,body,os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &models.Post{}, errors.New("Post Not Found")

}