package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateItem(client *dynamodb.Client) error {
	_, err := client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("23bcd59_student_table"),
		Key: map[string]types.AttributeValue{
			"studentId": &types.AttributeValueMemberS{Value: "1"},
		},
		UpdateExpression: aws.String("SET age = :a"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":a": &types.AttributeValueMemberN{Value: "22"},
		},
	})

	if err != nil {
		return err
	}
	return nil
}
