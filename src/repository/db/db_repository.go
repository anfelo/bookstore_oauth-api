package db

import (
	"fmt"

	"github.com/anfelo/bookstore_oauth-api/src/clients/mongodb"
	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/anfelo/bookstore_utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbRepository db repository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct{}

// NewRepository returns a new db repository
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	var accessToken *accesstoken.AccessToken
	client, ctx, cancel := mongodb.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("bookstore").Collection("oauth")
	result := collection.FindOne(ctx, bson.M{"access_token": id})
	if result.Err() != nil {
		return nil, errors.NewNotFoundError("access token not found")
	}
	err := result.Decode(&accessToken)
	if err != nil {
		fmt.Printf("error marshalling %v", err)
		return nil, errors.NewInternatServerError("database error")
	}

	return accessToken, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	client, ctx, cancel := mongodb.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("bookstore").Collection("oauth")
	_, err := collection.InsertOne(ctx, at)
	if err != nil {
		fmt.Println("error inserting access token", err)
		return errors.NewInternatServerError("database error")
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	client, ctx, cancel := mongodb.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	upsert := true
	opt := options.FindOneAndUpdateOptions{
		Upsert: &upsert,
	}
	collection := client.Database("bookstore").Collection("oauth")
	err := collection.FindOneAndUpdate(ctx, bson.M{"access_token": at.AccessToken}, at, &opt)
	if err != nil {
		fmt.Println("error updating access token expiration", err)
		return errors.NewInternatServerError("database error")
	}

	return nil
}
