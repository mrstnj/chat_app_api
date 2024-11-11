package services

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	e "github.com/mrstnj/chat_app_api/error"
	"github.com/mrstnj/chat_app_api/repository"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

func GetAllMessages(client _interface.DynamoDBClient) ([]byte, error) {
	message := repository.NewMessageRepository(client)
	out, err := message.FindByRoomId()
	if err != nil {
		return nil, e.DBError(err)
	}

	var room repository.MessageRoom
	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		return nil, err
	}

	messagesJSON, err := json.Marshal(room.Messages)
	if err != nil {
		return nil, err
	}

	return messagesJSON, nil
}

func PutMessages(client _interface.DynamoDBClient, req repository.Message) ([]byte, error) {
	message := repository.NewMessageRepository(client)

	out, err := message.FindByRoomId()
	if err != nil {
		return nil, e.DBError(err)
	}

	var room repository.MessageRoom
	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		return nil, err
	}

	messagesJSON, err := json.Marshal(room.Messages)
	if err != nil {
		return nil, err
	}

	room.Messages = append(room.Messages, repository.Message{
		Message:    req.Message,
		FromOthers: false,
		SendUser:   req.SendUser,
		SendTime:   time.Now(),
	})

	item, err := attributevalue.MarshalMap(room)
	if err != nil {
		return nil, err
	}

	if err := message.Update(item); err != nil {
		return nil, e.DBError(err)
	}

	return messagesJSON, nil
}
