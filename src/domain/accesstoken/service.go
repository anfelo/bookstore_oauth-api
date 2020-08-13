package accesstoken

import "github.com/anfelo/bookstore_oauth-api/src/utils/errors"

// Repository access token repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Service access token service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repo Repository
}

// NewService package function that return a new service
func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

// GetById method in chager of geting a service by id
func (s *service) GetByID(id string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
