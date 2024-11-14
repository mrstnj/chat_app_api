package connection

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/mrstnj/chat_app_api/error"
	"github.com/mrstnj/chat_app_api/handlers"
	"github.com/mrstnj/chat_app_api/repository"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
	"github.com/mrstnj/chat_app_api/services"
)

func ConnectHandler(client _interface.DynamoDBClient, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req repository.ConnectionId
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return handlers.ErrorResponse(e.InvalidValueError("params", err))
	}
	if err := services.Connect(client, req); err != nil {
		return handlers.ErrorResponse(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
