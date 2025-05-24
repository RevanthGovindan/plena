package services

import (
	"access-key-management/internal/database"
	"access-key-management/internal/models"
	"access-key-management/internal/stream"
	"access-key-management/pkg/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

func CreateNewAccessKeys() (models.AccessKeyResponse, error) {
	expTime := time.Now().Add(time.Hour * 2).Unix()
	var response = models.AccessKeyResponse{
		AccessKey: models.AccessKey{UserId: utils.GenerateRandom(),
			RateLimit: 100,
			Expiry:    expTime,
			Enabled:   true,
		},
	}
	var id = uuid.New()
	response.KeyId = id.String()

	err := database.GetDb().SaveAccessData(response.KeyId, response.AccessKey)
	if err != nil {
		return models.AccessKeyResponse{}, err
	}

	streamer := stream.GetStreamer()
	err = streamer.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_CREATED, Data: response})
	if err != nil {
		return models.AccessKeyResponse{}, err
	}
	return response, nil
}

func DeleteAccessKeys(keyId string) error {
	err := database.GetDb().DeleteAccessData(keyId)
	if err != nil {
		return err
	}
	streamer := stream.GetStreamer()
	var data = map[string]string{"keyId": keyId}
	err = streamer.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_DELETED, Data: data})
	return err
}

func UpdateAccessKeys(keyId string, keyData models.UpdateAccessKeyRequest) error {
	err := database.GetDb().UpdateAccessData(keyId, keyData)
	if err != nil {
		return err
	}
	streamer := stream.GetStreamer()
	var data = map[string]string{"keyId": keyId}
	err = streamer.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_UPDATED, Data: data})
	return err
}

func GetAllAccessKeys() (map[string]models.AccessKey, error) {
	data, err := database.GetDb().GetAllAccessData()
	if err != nil {
		return map[string]models.AccessKey{}, err
	}
	return data, nil
}

func GetDataByAccessKey(keyId string) (models.AccessKey, error) {
	data, exists := database.GetDb().GetAccessData(keyId)
	if !exists {
		return models.AccessKey{}, errors.New("not found")
	}
	if !data.Enabled {
		return models.AccessKey{}, errors.New("key disabled")
	}
	return data, nil
}

func DisableAccessKey(keyId string) error {
	err := database.GetDb().DisableAccessKey(keyId)
	return err
}
