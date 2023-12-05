package postcore_test

import (
	"context"
	"hexa-go/pkg/v1/datamodels"

	"github.com/stretchr/testify/mock"
)

type PostRepoMock struct {
	mock.Mock
}

func NewPostRepoMock() *PostRepoMock {
	return &PostRepoMock{}
}

func (repo *PostRepoMock) CreatePost(ctx context.Context, post datamodels.Post) error {
	args := repo.Called(ctx, post)
	return args.Error(0)

}

func (repo *PostRepoMock) GetPost() (results []datamodels.Post, err error) {
	args := repo.Called()

	return args.Get(0).([]datamodels.Post), args.Error(1)
}
