package accesstoken

import (
	"strings"
	"time"

	"github.com/anfelo/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenResquest main struct
type AccessTokenResquest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate method to validate an access token
func (r *AccessTokenResquest) Validate() *errors.RestErr {
	switch r.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type")
	}

	return nil
}

// AccessToken main struct
type AccessToken struct {
	AccessToken string `json:"access_token" bson:"access_token"`
	UserID      int64  `json:"user_id" bson:"user_id"`
	ClientID    int64  `json:"client_id" bson:"client_id"`
	Expires     int64  `json:"expires" bson:"expires"`
}

// GetNewAccessToken package function incharge of creating a new Access Token
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
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

// Generate method incharge of generating a valid token
func (at *AccessToken) Generate() {
	at.AccessToken = "some token"
}
