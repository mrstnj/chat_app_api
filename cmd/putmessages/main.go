package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mrstnj/chat_app_api/handlers/message"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

const domainName = "example.execute-api.region.amazonaws.com"
const stage = "prod"

func main() {
	lambda.Start(handleAPIGatewayRequest)
}

func handleAPIGatewayRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client, wsClient, err := initDynamoDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       `{"message":"database operation failed"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}
	return message.PutMessagesHandler(client, wsClient, request)
}

func initDynamoDB() (_interface.DynamoDBClient, _interface.WebSocketClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		return nil, nil, err
	}

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://dynamodb-local:8000/")
	})

	endpoint := fmt.Sprintf("https://%s/%s", domainName, stage)
	wsClient := apigatewaymanagementapi.NewFromConfig(cfg, func(o *apigatewaymanagementapi.Options) {
		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(endpoint)
	})

	return client, wsClient, nil
}
