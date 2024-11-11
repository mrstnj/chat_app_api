package message

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/mrstnj/chat_app_api/error"
	"github.com/mrstnj/chat_app_api/handlers"
	"github.com/mrstnj/chat_app_api/repository"
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
	var req repository.Message
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return handlers.ErrorResponse(e.InvalidValueError("params", err))
	}
	messagesJSON, err := services.PutMessages(client, req)
	if err != nil {
		return handlers.ErrorResponse(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}
