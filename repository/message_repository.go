package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	e "github.com/mrstnj/chat_app_api/error"
	_interface "github.com/mrstnj/chat_app_api/repository/interface"
)

type Message struct {
	Message     string    `json:"message"`
	FromChatGPT bool      `json:"from_others"`
	SendUser    string    `json:"send_user"`
	SendTime    time.Time `json:"send_time"`
}

type MessageRoom struct {
	RoomID   int       `json:"room_id"`
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

func (r *MessageRepository) Update(item map[string]types.AttributeValue) error {
	if _, err := r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("message_rooms"),
		Item: map[string]types.AttributeValue{
			"room_id":  &types.AttributeValueMemberN{Value: "1"},
			"messages": item["Messages"],
		},
	}); err != nil {
		return e.DBError(err)
	}

	return nil
}
