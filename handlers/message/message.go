package message

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
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

func PutMessagesHandler(client _interface.DynamoDBClient, wsClient _interface.WebSocketClient, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req repository.Message
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return handlers.ErrorResponse(e.InvalidValueError("params", err))
	}
	res, err := services.PutMessages(client, req)
	if err != nil {
		return handlers.ErrorResponse(err)
	}

	messageData, err := json.Marshal(res.Messages)
	if err != nil {
		return handlers.ErrorResponse(err)
	}

	for _, connectionId := range res.ConnectionIds {
		if _, err := wsClient.PostToConnection(context.TODO(), &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(strconv.Itoa(connectionId)),
			Data:         messageData,
		}); err != nil {
			return handlers.ErrorResponse(err)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
