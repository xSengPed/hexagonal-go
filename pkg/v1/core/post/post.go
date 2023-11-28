package postcore

import "hexa-go/pkg/v1/core/post/repository"

type service struct {
	db repository.IPostRepositoryDB
}

func NewService(db repository.IPostRepositoryDB) *service {
	return &service{db}
}
