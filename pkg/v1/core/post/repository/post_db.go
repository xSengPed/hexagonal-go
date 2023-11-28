package repository

import (
	"context"
	"hexa-go/pkg/v1/datamodels"

	"go.mongodb.org/mongo-driver/mongo"
)

type postRepositoryDB struct {
	collection *mongo.Collection
}

type IPostRepositoryDB interface {
	//
	CreatePost(ctx context.Context, post datamodels.Post) error
	GetPost() (results []datamodels.Post, err error)
}

func NewPostRepositoryDB(mongoClient *mongo.Client) *postRepositoryDB {
	return &postRepositoryDB{mongoClient.Database("my-app").Collection("post")}
}

func (repo *postRepositoryDB) CreatePost(ctx context.Context, post datamodels.Post) error {
	_, err := repo.collection.InsertOne(ctx, post)

	if err != nil {
		return err
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
