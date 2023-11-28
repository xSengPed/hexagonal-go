package commentrepo

import (
	"context"
	"hexa-go/pkg/v1/datamodels"

	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepositoryDB struct {
	collection *mongo.Collection
}

type ICommentRepositoryDB interface {
	CreateComment(ctx context.Context, comment datamodels.Comment) error
}

func NewCommentRepositoryDB(mongoClient *mongo.Client) *CommentRepositoryDB {
	return &CommentRepositoryDB{mongoClient.Database("my-app").Collection("comment")}
}

func (repo *CommentRepositoryDB) CreateComment(ctx context.Context, post datamodels.Comment) error {
	_, err := repo.collection.InsertOne(ctx, post)

	if err != nil {
		return err
	}
	return nil
}
