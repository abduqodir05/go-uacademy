package controller

import (
	"go-uacademy/config"
	"go-uacademy/storage"
)

type Controller struct {
	cfg   *config.Config
	store *storage.Store
}

// Controller Constructor
func NewController(cfg *config.Config, store *storage.Store) *Controller {
	return &Controller{
		cfg:   cfg,
		store: store,
	}
}
