package services

import (
	"strings"

	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/anfelo/bookstore_oauth-api/src/domain/users"
	"github.com/anfelo/bookstore_utils/errors"
)

// Repository access token repository interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

// RestRepository sets repository interace
type RestRepository interface {
	Login(string, string) (*users.User, *errors.RestErr)
}

// Service access token service interface
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessTokenResquest) (*accesstoken.AccessToken, *errors.RestErr)
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type service struct {
	dbRepo        Repository
	restUsersRepo RestRepository
}

// NewService package function that return a new service
func NewService(dbRepo Repository, restUsersRepo RestRepository) Service {
	return &service{
		dbRepo,
		restUsersRepo,
	}
}

// GetById method in chager of geting a service by id
func (s *service) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	accessTokenID := strings.TrimSpace(id)
	if accessTokenID == "" {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request accesstoken.AccessTokenResquest) (*accesstoken.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// TODO: Create login logic for grant_type=client_credentials

	// Login logic for grant_type=password
	user, err := s.restUsersRepo.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := accesstoken.GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
