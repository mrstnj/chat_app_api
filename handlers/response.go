package handlers

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	e "github.com/mrstnj/chat_app_api/error"
)

func ErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	if appErr, ok := err.(*e.AppError); ok {
		code := ResponseCode(appErr)
		return events.APIGatewayProxyResponse{
			StatusCode: code,
			Body:       fmt.Sprintf(`{"messages":"%s"}`, appErr.Msg),
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       `{"messages":"an unexpected error occurred"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func SuccessResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"messages":"request completed successfully"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func ResponseCode(appErr *e.AppError) int {
	switch appErr.ErrType {
	case e.InvalidValueErr:
		return 401
	case e.NotFoundErr:
		return 404
	case e.RequiredValueErr:
		return 422
	case e.APIErr:
		return 502
	case e.DBErr, e.StorageErr, e.EmailErr:
		return 503
	default:
		return 500
	}
}
