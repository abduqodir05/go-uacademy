package jsondb

import (
	"go-uacademy/config"
	"go-uacademy/storage"
	"os"
)

type Store struct {
	user *userRepo
	post *postRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	},nil

	
}
func NewPostFileJson(cfg *config.Config) (storage.StorageI, error) {

	postFile, err := os.Open(cfg.Path + cfg.PostFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg. PostFileName, postFile),
	},nil

	
}
func (s *Store) CloseDB() {
	s.user.file.Close()
	s.post.file.Close()
}

func (s *Store) User() storage.UserRepoI{
	return s.user
}
func (s *Store) Post() storage.PostReopI {
	return s.post
}