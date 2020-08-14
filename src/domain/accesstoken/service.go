package accesstoken

import (
	"strings"

	"github.com/anfelo/bookstore_oauth-api/src/utils/errors"
)

// Repository access token repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// Service access token service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
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
	accessTokenID := strings.TrimSpace(id)
	if accessTokenID == "" {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repo.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateExpirationTime(at)
}
