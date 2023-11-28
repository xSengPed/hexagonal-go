package postcore

import "hexa-go/pkg/v1/datamodels"

type CreatePostFunc func(post datamodels.Post) error

func (s *service) CreatePostFunc(post datamodels.Post) error {

	err := s.db.CreatePost(post)

	if err != nil {
		return err
	}

	return nil
}
