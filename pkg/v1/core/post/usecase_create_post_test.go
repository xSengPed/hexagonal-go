package postcore_test

import (
	"context"
	"errors"
	postcore "hexa-go/pkg/v1/core/post"
	"hexa-go/pkg/v1/datamodels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	ctx := context.Background()
	t.Run("Test Create Post Success !", func(t *testing.T) {
		// 1. Arrange
		postModel := datamodels.Post{
			Title:   "test",
			Content: "test content",
		}
		mockRepo := NewPostRepoMock()
		mockRepo.On("CreatePost", ctx, postModel).Return(nil)

		postService := postcore.NewService(mockRepo)

		// 2. Act

		err := postService.CreatePostFunc(ctx, postModel)

		// 3. Assert

		assert.Nil(t, err)
	})

	t.Run("Test Create Post Fail !", func(t *testing.T) {
		// 1. Arrange
		postModel := datamodels.Post{
			Title:   "test",
			Content: "test content",
		}
		mockRepo := NewPostRepoMock()
		mockRepo.On("CreatePost", ctx, postModel).Return(errors.New(""))

		postService := postcore.NewService(mockRepo)

		// 2. Act

		err := postService.CreatePostFunc(ctx, postModel)

		// 3. Assert

		assert.NotNil(t, err)
	})

}
