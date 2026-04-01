package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutItem(client *dynamodb.Client) error {
	_, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("23bcd59_student_table"),
		Item: map[string]types.AttributeValue{
			"studentId": &types.AttributeValueMemberS{Value: "23bcd59"},
			"name":      &types.AttributeValueMemberS{Value: "Harissh"},
			"age":       &types.AttributeValueMemberN{Value: "21"},
		},
	})

	return err
}
