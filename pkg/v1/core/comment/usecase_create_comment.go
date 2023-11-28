package commentcore

import (
	"context"
	"hexa-go/pkg/v1/datamodels"
)

type CreateComment func(ctx context.Context, comment datamodels.Comment) error

func (s *service) CreateComment(ctx context.Context, comment datamodels.Comment) error {

	err := s.db.CreateComment(ctx, comment)

	if err != nil {
		return err
	}

	return nil
}
