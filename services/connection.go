package services

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	e "github.com/mrstnj/chat_app_api/error"
	"github.com/mrstnj/chat_app_api/repository"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

func Connect(client _interface.DynamoDBClient, req repository.ConnectionId) error {
	var room repository.MessageRoom
	message := repository.NewMessageRepository(client)

	out, err := message.FindByRoomId()
	if err != nil {
		return e.DBError(err)
	}

	if err := attributevalue.UnmarshalMap(out.Item, &room); err != nil {
		return err
	}

	if _, err = json.Marshal(room.Messages); err != nil {
		return err
	}

	room.ConnectionIds = append(room.ConnectionIds, req.ConnectionId)

	item, err := attributevalue.MarshalMap(room)
	if err != nil {
		return err
	}

	if err := message.UpdateConnectionIds(item); err != nil {
		return e.DBError(err)
	}

	return nil
}
