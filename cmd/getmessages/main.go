package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mrstnj/chat_app_api/handlers/messages"
)

func main() {
	lambda.Start(messages.GetAllMessagesHandler)
}
