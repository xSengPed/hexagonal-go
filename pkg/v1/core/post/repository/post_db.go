package repository

import (
	"errors"
	"hexa-go/pkg/v1/datamodels"
)

type postRepositoryDB struct {
	//
}

type IPostRepositoryDB interface {
	//
	CreatePost(post datamodels.Post) error
	GetPost() (results []datamodels.Post, err error)
}

func NewPostRepositoryDB() *postRepositoryDB {
	return &postRepositoryDB{}
}

func (repo *postRepositoryDB) CreatePost(post datamodels.Post) error {
	if post.Title == "rose" {
		return errors.New("Fatal Error : Rosiier Too Beautiful")
	}
	return nil
}

func (repo *postRepositoryDB) GetPost() (results []datamodels.Post, err error) {
	results = []datamodels.Post{
		{
			Title:   "Test 1",
			Content: "...",
		},
		{
			Title:   "Test 2",
			Content: "...",
		},
	}
	return results, nil
}
