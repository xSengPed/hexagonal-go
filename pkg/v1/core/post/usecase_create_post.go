package postcore

import (
	"context"
	"hexa-go/pkg/v1/datamodels"
)

type CreatePostFunc func(ctx context.Context, post datamodels.Post) error

func (s *service) CreatePostFunc(ctx context.Context, post datamodels.Post) error {

	err := s.db.CreatePost(ctx, post)

	if err != nil {
		return err
	}

	return nil
}
