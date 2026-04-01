package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetItem(client *dynamodb.Client) error {
	res, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("23bcd59_student_table"),
		Key: map[string]types.AttributeValue{
			"studentId": &types.AttributeValueMemberS{Value: "1"},
		},
	})

	if err != nil {
		return err
	}

	fmt.Println(res.Item)
	return nil
}
