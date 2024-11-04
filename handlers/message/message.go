package message

import (
	"github.com/aws/aws-lambda-go/events"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
	"github.com/mrstnj/chat_app_api/services"
	"log"
)

func GetAllMessagesHandler(client _interface.DynamoDBClient, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	messagesJSON, err := services.GetAllMessages(client)

	if err != nil {
		log.Fatalf("failed to marshal messages, %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(messagesJSON),
		StatusCode: 200,
	}, nil
}
