package db

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/anil1226/go-employee-dynamo/internal/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) GetUser(ctx context.Context, user models.User) (models.User, error) {

	result, err := d.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(containerUser),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(user.Name),
			},
		},
	})
	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		return models.User{}, nil
	}

	if result.Item == nil {
		msg := "Could not find '" + user.Name + "'"
		return models.User{}, errors.New(msg)
	}

	itemResponseBody := models.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &itemResponseBody)
	if err != nil {
		log.Printf("Failed to unmarshal Record: %v", err)
		return models.User{}, nil
	}

	if len(strings.Trim(itemResponseBody.Name, "")) == 0 {
		return models.User{}, errors.New("user does not exist")
	}

	if bcrypt.CompareHashAndPassword([]byte(itemResponseBody.Password), []byte(user.Password)) == nil {
		return itemResponseBody, nil
	}
	return models.User{}, errors.New("password does not match")

}

func (d *Database) CreateUser(ctx context.Context, user models.User) error {

	user.ID = uuid.NewV4().String()
	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Printf("Got error marshalling item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(containerUser),
	}

	_, err = d.Client.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return err
	}

	return nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
