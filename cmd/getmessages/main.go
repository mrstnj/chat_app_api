package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mrstnj/chat_app_api/handlers/messages"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

func main() {
	lambda.Start(handleAPIGatewayRequest)
}

func handleAPIGatewayRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client, _ := initDynamoDB()
	return messages.GetAllMessagesHandler(client, request)
}

func initDynamoDB() (_interface.DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://dynamodb-local:8000/")
	})

	return client, nil
}
