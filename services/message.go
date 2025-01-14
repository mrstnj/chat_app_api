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

func PutMessages(client _interface.DynamoDBClient, req repository.Message) (repository.MessageRoom, error) {
	var room repository.MessageRoom
	message := repository.NewMessageRepository(client)

	out, err := message.FindByRoomId()
	if err != nil {
		return room, e.DBError(err)
	}

	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		return room, err
	}

	if _, err = json.Marshal(room.Messages); err != nil {
		return room, err
	}

	room.Messages = append(room.Messages, repository.Message{
		Message:     req.Message,
		FromChatGPT: false,
		SendUser:    req.SendUser,
		SendTime:    time.Now(),
	})

	item, err := attributevalue.MarshalMap(room)
	if err != nil {
		return room, err
	}

	if err := message.UpdateMessages(item); err != nil {
		return room, e.DBError(err)
	}

	return room, nil
}
