package connection

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

func TestConnectHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name           string
		mockSetup      func(mockDynamoDB *_mock.MockDynamoDBClient)
		expectedStatus int
	}{
		{
			name: "success case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"room_id": &types.AttributeValueMemberN{Value: "1"},
						"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"message":     &types.AttributeValueMemberS{Value: "Hello"},
								"fromChatCPT": &types.AttributeValueMemberBOOL{Value: false},
								"sendUser":    &types.AttributeValueMemberS{Value: "Ken"},
								"sendTime":    &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
							}},
						}},
						"connection_ids": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "1"},
						}},
					},
				}, nil)
				mockDynamoDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "DynamoDB GET error case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("DynamoDB error"))
			},
			expectedStatus: 503,
		},
		{
			name: "DynamoDB PUT error case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"room_id": &types.AttributeValueMemberN{Value: "1"},
						"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"message":     &types.AttributeValueMemberS{Value: "Hello"},
								"fromChatCPT": &types.AttributeValueMemberBOOL{Value: false},
								"sendUser":    &types.AttributeValueMemberS{Value: "Ken"},
								"sendTime":    &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
							}},
						}},
						"connection_ids": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "1"},
						}},
					},
				}, nil)
				mockDynamoDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("DynamoDB error"))
			},
			expectedStatus: 503,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockDynamoDB := _mock.NewMockDynamoDBClient(ctrl)

			testCase.mockSetup(mockDynamoDB)

			request := events.APIGatewayProxyRequest{
				Body: "{\"connection_id\": 2}",
			}
			response, err := ConnectHandler(mockDynamoDB, request)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedStatus, response.StatusCode)
		})
	}
}

func TestDisconnectHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name           string
		mockSetup      func(mockDynamoDB *_mock.MockDynamoDBClient)
		expectedStatus int
	}{
		{
			name: "success case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"room_id": &types.AttributeValueMemberN{Value: "1"},
						"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"message":     &types.AttributeValueMemberS{Value: "Hello"},
								"fromChatCPT": &types.AttributeValueMemberBOOL{Value: false},
								"sendUser":    &types.AttributeValueMemberS{Value: "Ken"},
								"sendTime":    &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
							}},
						}},
						"connection_ids": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "1"},
						}},
					},
				}, nil)
				mockDynamoDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "DynamoDB GET error case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("DynamoDB error"))
			},
			expectedStatus: 503,
		},
		{
			name: "DynamoDB PUT error case",
			mockSetup: func(mockDynamoDB *_mock.MockDynamoDBClient) {
				mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"room_id": &types.AttributeValueMemberN{Value: "1"},
						"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
								"message":     &types.AttributeValueMemberS{Value: "Hello"},
								"fromChatCPT": &types.AttributeValueMemberBOOL{Value: false},
								"sendUser":    &types.AttributeValueMemberS{Value: "Ken"},
								"sendTime":    &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
							}},
						}},
						"connection_ids": &types.AttributeValueMemberL{Value: []types.AttributeValue{
							&types.AttributeValueMemberN{Value: "1"},
						}},
					},
				}, nil)
				mockDynamoDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("DynamoDB error"))
			},
			expectedStatus: 503,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockDynamoDB := _mock.NewMockDynamoDBClient(ctrl)

			testCase.mockSetup(mockDynamoDB)

			request := events.APIGatewayProxyRequest{
				Body: "{\"connection_id\": 2}",
			}
			response, err := DisconnectHandler(mockDynamoDB, request)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedStatus, response.StatusCode)
		})
	}
}
