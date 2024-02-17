package db

import (
	"log"

	"github.com/anil1226/go-employee-dynamo/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database struct {
	Client *dynamodb.DynamoDB
}

var (
	// cosmosDbEndpoint = config.GetEnvKey("cosmosDbEndpoint")
	// cosmosDbKey      = config.GetEnvKey("cosmosDbKey")
	//dbName        = config.GetEnvKey("dbName")
	containerEmp  = config.GetEnvKey("containerEmp")
	containerUser = config.GetEnvKey("containerUser")
)

func NewDatabase() (*Database, error) {

	// Create a new AWS session with shared credentials from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),                         // Specify your AWS region
		Credentials: credentials.NewSharedCredentials("", "default"), // Use the default credentials profile
	})
	if err != nil {
		log.Fatal("Error creating session:", err)

	}

	// Create a DynamoDB client
	svc := dynamodb.New(sess)

	return &Database{
		Client: svc,
	}, nil
}
