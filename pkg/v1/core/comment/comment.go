package commentcore

import commentrepo "hexa-go/pkg/v1/core/comment/repository"

type service struct {
	db commentrepo.ICommentRepositoryDB
}

func NewService(db commentrepo.ICommentRepositoryDB) *service {
	return &service{db}
}
