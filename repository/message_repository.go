package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	e "github.com/mrstnj/chat_app_api/error"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

type Message struct {
	Message    string `json:"message"`
	FromOthers bool   `json:"from_others"`
	SendTime   string `json:"send_time"`
}

type MessageRoom struct {
	Messages []Message `json:"messages"`
}

type MessageRepository struct {
	client _interface.DynamoDBClient
}

func NewMessageRepository(client _interface.DynamoDBClient) *MessageRepository {
	return &MessageRepository{
		client: client,
	}
}

func (r *MessageRepository) FindByRoomId() (*dynamodb.GetItemOutput, error) {
	out, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("message_rooms"),
		Key: map[string]types.AttributeValue{
			"room_id": &types.AttributeValueMemberN{Value: "1"},
		},
	})
	if err != nil {
		return nil, e.DBError(err)
	}

	return out, nil
}
