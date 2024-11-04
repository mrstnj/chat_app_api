package message

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/mrstnj/chat_app_api/repository"
	_mock "github.com/mrstnj/chat_app_api/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMessagesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDynamoDB := _mock.NewMockDynamoDBClient(ctrl)

	mockDynamoDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(&dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"room_id": &types.AttributeValueMemberN{Value: "1"},
			"messages": &types.AttributeValueMemberL{Value: []types.AttributeValue{
				&types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
					"message":    &types.AttributeValueMemberS{Value: "わんわん"},
					"fromOthers": &types.AttributeValueMemberBOOL{Value: true},
					"sendTime":   &types.AttributeValueMemberS{Value: "2023-11-03T00:00:00Z"},
				}},
			}},
		},
	}, nil)

	request := events.APIGatewayProxyRequest{}
	response, err := GetAllMessagesHandler(mockDynamoDB, request)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	var messages []repository.Message
	err = json.Unmarshal([]byte(response.Body), &messages)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, "わんわん", messages[0].Message)
	assert.Equal(t, true, messages[0].FromOthers)
	assert.Equal(t, "2023-11-03T00:00:00Z", messages[0].SendTime)
}
