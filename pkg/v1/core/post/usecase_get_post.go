package postcore

import "hexa-go/pkg/v1/datamodels"

type GetPostFunc func() (result []datamodels.Post, err error)

func (s *service) GetPostFunc() (result []datamodels.Post, err error) {
	// connect to repo
	results, err := s.db.GetPost()

	if err != nil {
		return results, err
	}
	return results, nil
}
