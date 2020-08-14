package accesstoken

import (
	"strings"
	"time"

	"github.com/anfelo/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken main struct
type AccessToken struct {
	AccessToken string `json:"access_token" bson:"access_token"`
	UserID      int64  `json:"user_id" bson:"user_id"`
	ClientID    int64  `json:"client_id" bson:"client_id"`
	Expires     int64  `json:"expires" bson:"expires"`
}

// GetNewAccessToken package function incharge of creating a new Access Token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired method that tells if an Access Token is expired
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

// Validate method to validate an access token
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
