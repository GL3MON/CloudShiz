package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteItem(client *dynamodb.Client) error {
	_, err := client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("23bcd59_student_table"),
		Key: map[string]types.AttributeValue{
			"studentId": &types.AttributeValueMemberS{Value: "1"},
		},
	})

	return err
}
