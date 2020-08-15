package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"

	"github.com/anfelo/bookstore_oauth-api/src/domain/users"
	"github.com/anfelo/bookstore_utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

// UsersRepository rest users repository interface
type UsersRepository interface {
	Login(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

// NewRepository returns a new rest user repository
func NewRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) Login(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternatServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternatServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternatServerError("error when trying to unmarshal user response")
	}
	return &user, nil
}
