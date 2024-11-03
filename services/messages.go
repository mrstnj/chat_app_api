package services

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/mrstnj/chat_app_api/repository"
)

type Message struct {
	messageRepo *repository.MessageRepository
}

func NewMessage(messageRepo *repository.MessageRepository) *Message {
	return &Message{
		messageRepo: messageRepo,
	}
}

func (s *Message) GetAllMessages() ([]byte, error) {
	out, err := s.messageRepo.GetAllMessages()
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
