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

type AccessKeyServices struct {
	db     database.Database
	stream stream.Stream
}

func NewAccessKeyService() AccessKeyServices {
	return AccessKeyServices{
		db:     database.GetDb(),
		stream: stream.GetStreamer(),
	}
}

func (f *AccessKeyServices) CreateNewAccessKeys() (models.AccessKeyResponse, error) {
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

	err := f.db.SaveAccessData(response.KeyId, response.AccessKey)
	if err != nil {
		return models.AccessKeyResponse{}, err
	}

	err = f.stream.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_CREATED, Data: response})
	if err != nil {
		return models.AccessKeyResponse{}, err
	}
	return response, nil
}

func (f *AccessKeyServices) DeleteAccessKeys(keyId string) error {
	err := f.db.DeleteAccessData(keyId)
	if err != nil {
		return err
	}
	var data = map[string]string{"keyId": keyId}
	err = f.stream.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_DELETED, Data: data})
	return err
}

func (f *AccessKeyServices) UpdateAccessKeys(keyId string, keyData models.UpdateAccessKeyRequest) error {
	data, err := f.db.UpdateAccessData(keyId, keyData)
	if err != nil {
		return err
	}
	var response models.AccessKeyResponse = models.AccessKeyResponse{
		AccessKey: data,
		KeyId:     keyId,
	}
	err = f.stream.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_UPDATED, Data: response})
	return err
}

func (f *AccessKeyServices) GetAllAccessKeys() (map[string]models.AccessKey, error) {
	data, err := f.db.GetAllAccessData()
	if err != nil {
		return map[string]models.AccessKey{}, err
	}
	return data, nil
}

func (f *AccessKeyServices) GetDataByAccessKey(keyId string) (models.AccessKey, error) {
	data, exists := f.db.GetAccessData(keyId)
	if !exists {
		return models.AccessKey{}, errors.New("not found")
	}
	if !data.Enabled {
		return models.AccessKey{}, errors.New("key disabled")
	}
	return data, nil
}

func (f *AccessKeyServices) DisableAccessKey(keyId string) error {
	err := f.db.DisableAccessKey(keyId)
	if err != nil {
		return err
	}
	var data = map[string]string{"keyId": keyId}
	err = f.stream.Publish(utils.PUBLISH_TOPIC, models.EventMessage{Event: utils.ACCESSKEY_DISABLED, Data: data})
	return err
}
