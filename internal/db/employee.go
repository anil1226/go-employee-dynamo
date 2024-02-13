package db

import (
	"context"
	"errors"
	"log"

	"github.com/anil1226/go-employee-dynamo/internal/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

func (d *Database) GetEmployee(ctx context.Context, id string) (models.Employee, error) {
	result, err := d.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(containerEmp),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		return models.Employee{}, nil
	}

	if result.Item == nil {
		msg := "Could not find '" + id + "'"
		return models.Employee{}, errors.New(msg)
	}

	itemResponseBody := models.Employee{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &itemResponseBody)
	if err != nil {
		log.Printf("Failed to unmarshal Record: %v", err)
		return models.Employee{}, nil
	}
	return itemResponseBody, nil

}

func (d *Database) CreateEmployee(ctx context.Context, emp models.Employee) error {
	emp.ID = uuid.NewV4().String()

	av, err := dynamodbattribute.MarshalMap(emp)
	if err != nil {
		log.Printf("Got error marshalling item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(containerEmp),
	}

	_, err = d.Client.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return err
	}

	return nil
}

func (d *Database) UpdateEmployee(ctx context.Context, emp models.Employee) error {
	av, err := dynamodbattribute.MarshalMap(emp)
	if err != nil {
		log.Printf("Got error marshalling new movie item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(containerEmp),
	}

	_, err = d.Client.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
		return err
	}
	return nil
}

func (d *Database) DeleteEmployee(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(containerEmp),
	}

	_, err := d.Client.DeleteItem(input)
	if err != nil {
		log.Printf("Got error calling DeleteItem: %s", err)
		return err
	}
	return nil
}

// func (d *Database) GetContainerClient(name string) (*azcosmos.ContainerClient, error) {
// 	return d.Client.NewContainer(name)

// }
