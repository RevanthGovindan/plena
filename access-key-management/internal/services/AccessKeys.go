package services

import (
	"access-key-management/internal/models"
	"access-key-management/internal/stream"
	"access-key-management/pkg/utils"
	"time"

	"github.com/google/uuid"
)

func CreateNewAccessKeys() (models.AccessKeyResponse, error) {
	expTime := time.Now().Add(time.Hour * 2).Unix()
	var response = models.AccessKeyResponse{
		AccessKey: models.AccessKey{UserId: utils.GenerateRandom(),
			RateLimit: 100,
			Expiry:    expTime,
		},
	}
	var id = uuid.New()
	response.KeyId = id.String()
	streamer := stream.GetStreamer()
	err := streamer.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_CREATED, Data: response})
	if err != nil {
		return models.AccessKeyResponse{}, err
	}
	return response, nil
}

func DeleteAccessKeys(accessId string) error {
	streamer := stream.GetStreamer()
	var data = map[string]string{"keyId": accessId}
	err := streamer.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_DELETED, Data: data})
	return err
}
