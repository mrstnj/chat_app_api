package message

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/mrstnj/chat_app_api/handlers"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
	"github.com/mrstnj/chat_app_api/services"
)

func GetAllMessagesHandler(client _interface.DynamoDBClient, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	messagesJSON, err := services.GetAllMessages(client)
	if err != nil {
		return handlers.ErrorResponse(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}

func PutMessagesHandler(client _interface.DynamoDBClient, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	messagesJSON, err := services.PutMessages(client)
	if err != nil {
		return handlers.ErrorResponse(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}
