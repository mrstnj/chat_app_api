package messages

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mrstnj/chat_app_api/repository"
	"github.com/mrstnj/chat_app_api/services"
)

func initDynamoDB() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://dynamodb-local:8000/")
	})

	return client, nil
}

func GetAllMessagesHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client, err := initDynamoDB()
	if err != nil {
		log.Fatalf("failed to connect DynamoDB, %v", err)
	}

	message := services.NewMessage(repository.NewMessageRepository(client))
	messagesJSON, err := message.GetAllMessages()
	if err != nil {
		log.Fatalf("failed to connect DynamoDB, %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}
