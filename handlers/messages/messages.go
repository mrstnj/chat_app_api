package messages

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Message struct {
	Message  string `json:"message"`
	SendUser string `json:"send_user"`
}

type MessageRoom struct {
	Messages []Message `json:"messages"`
}

func GetAllMessagesHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://dynamodb-local:8000/")
	})

	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("message_rooms"),
		Key: map[string]types.AttributeValue{
			"room_id": &types.AttributeValueMemberN{Value: "1"},
		},
	})
	if err != nil {
		log.Fatalf("failed to get item, %v", err)
	}

	var room MessageRoom
	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		log.Fatalf("failed to unmarshal result item, %v", err)
	}

	messagesJSON, err := json.Marshal(room.Messages)
	if err != nil {
		log.Fatalf("failed to marshal messages, %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}
