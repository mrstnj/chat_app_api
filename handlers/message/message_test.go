package message

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	_mock "github.com/mrstnj/chat_app_api/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessagesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name           string
		mockSetup      func(mockDynamoDB *_mock.MockDynamoDBClient)
		expectedBody   string
		expectedStatus int
		expectError    bool
	}{
		{
			name: "success case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"room_id": &types.AttributeValueMemberN{Value: "1"},
						"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"message":    &types.AttributeValueMemberS{Value: "Hello"},
								"fromChatCPT": &types.AttributeValueMemberBOOL{Value: false},
								"sendUser":   &types.AttributeValueMemberS{Value: "Ken"},
								"sendTime":   &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
							}},
						}},
					},
				}, nil)
			},
			expectedBody:   `[{"message":"Hello","from_others":false,"send_user":"Ken","send_time":"2023-11-03T00:00:00Z"}]`,
			expectedStatus: 200,
		},
		{
			name: "DynamoDB error case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("DynamoDB error"))
			},
			expectedBody:   `{"messages":"database operation failed"}`,
			expectedStatus: 503,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockDynamoDB := _mock.NewMockDynamoDBClient(ctrl)
			testCase.mockSetup(mockDynamoDB)

			request := events.APIGatewayProxyRequest{}
			response, err := GetAllMessagesHandler(mockDynamoDB, request)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedStatus, response.StatusCode)
			assert.JSONEq(t, testCase.expectedBody, response.Body)
		})
	}
}
