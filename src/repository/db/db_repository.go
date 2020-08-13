package db

import (
	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/anfelo/bookstore_oauth-api/src/utils/errors"
)

// DbRepository db repository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

// NewRepository returns a new db repository
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternatServerError("database connection not yet implemented")
}
