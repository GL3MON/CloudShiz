package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTable(client *dynamodb.Client) error {
	_, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("23bcd59_student_table"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("StudentId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("StudentId"),
				KeyType:       types.KeyTypeHash,
			},
		},

		BillingMode: types.BillingModePayPerRequest,
	})

	return err
}
