package helloworld

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HelloWorldHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}
