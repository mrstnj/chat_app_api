package services

import (
	"encoding/json"
	
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/mrstnj/chat_app_api/repository"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
	"log"
)

func GetAllMessages(client _interface.DynamoDBClient) ([]byte, error) {
	message := repository.NewMessageRepository(client)
	out, err := message.FindByRoomId()
	if err != nil {
		log.Fatalf("failed to unmarshal result item, %v", err)
	}

	var room repository.MessageRoom
	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		log.Fatalf("failed to unmarshal result item, %v", err)
	}

	messagesJSON, err := json.Marshal(room.Messages)
	if err != nil {
		log.Fatalf("failed to marshal messages, %v", err)
	}

	return messagesJSON, nil
}
