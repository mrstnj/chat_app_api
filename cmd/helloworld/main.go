package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mrstnj/chat_app_api/handlers/helloworld"
)

func main() {
	lambda.Start(helloworld.HelloWorldHandler)
}
