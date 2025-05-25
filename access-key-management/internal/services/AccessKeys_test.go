package services

import (
	"access-key-management/internal/models"
	"access-key-management/pkg/utils"
	"access-key-management/testutils"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewAccessKeys_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	mockDb.On("SaveAccessData", mock.Anything, mock.Anything).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(nil)

	response, err := service.CreateNewAccessKeys()

	assert.NoError(t, err)
	assert.NotEmpty(t, response.KeyId)
	assert.NotEmpty(t, response.AccessKey.UserId)
	assert.Equal(t, 100, response.AccessKey.RateLimit)
	assert.True(t, response.AccessKey.Enabled)
}

func TestCreateNewAccessKeys_SaveAccessDataError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	mockDb.On("SaveAccessData", mock.Anything, mock.Anything).Return(errors.New("db error"))
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(nil)

	_, err := service.CreateNewAccessKeys()

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

func TestCreateNewAccessKeys_PublishError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	mockDb.On("SaveAccessData", mock.Anything, mock.Anything).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(errors.New("publish error"))

	_, err := service.CreateNewAccessKeys()

	assert.Error(t, err)
	assert.Equal(t, "publish error", err.Error())
	mockDb.AssertCalled(t, "SaveAccessData", mock.Anything, mock.Anything)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}
func TestUpdateAccessKeys_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"
	updateRequest := models.UpdateAccessKeyRequest{
		RateLimit: 200,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
	}
	updatedAccessKey := models.AccessKey{
		UserId:    1,
		RateLimit: 200,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
		Enabled:   false,
	}

	mockDb.On("UpdateAccessData", keyId, updateRequest).Return(updatedAccessKey, nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(nil)

	err := service.UpdateAccessKeys(keyId, updateRequest)

	assert.NoError(t, err)
	mockDb.AssertCalled(t, "UpdateAccessData", keyId, updateRequest)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestUpdateAccessKeys_UpdateAccessDataError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"
	updateRequest := models.UpdateAccessKeyRequest{
		RateLimit: 200,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
	}

	mockDb.On("UpdateAccessData", keyId, updateRequest).Return(models.AccessKey{}, errors.New("update error"))

	err := service.UpdateAccessKeys(keyId, updateRequest)

	assert.Error(t, err)
	assert.Equal(t, "update error", err.Error())
	mockDb.AssertCalled(t, "UpdateAccessData", keyId, updateRequest)
	mockStream.AssertNotCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestUpdateAccessKeys_PublishError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"
	updateRequest := models.UpdateAccessKeyRequest{
		RateLimit: 200,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
	}
	updatedAccessKey := models.AccessKey{
		UserId:    1,
		RateLimit: 200,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
		Enabled:   false,
	}

	mockDb.On("UpdateAccessData", keyId, updateRequest).Return(updatedAccessKey, nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(errors.New("publish error"))

	err := service.UpdateAccessKeys(keyId, updateRequest)

	assert.Error(t, err)
	assert.Equal(t, "publish error", err.Error())
	mockDb.AssertCalled(t, "UpdateAccessData", keyId, updateRequest)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestGetAllAccessKeys_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	mockData := map[string]models.AccessKey{
		"key1": {
			UserId:    1,
			RateLimit: 100,
			Expiry:    time.Now().Add(time.Hour * 2).Unix(),
			Enabled:   true,
		},
		"key2": {
			UserId:    2,
			RateLimit: 200,
			Expiry:    time.Now().Add(time.Hour * 3).Unix(),
			Enabled:   false,
		},
	}

	mockDb.On("GetAllAccessData").Return(mockData, nil)

	data, err := service.GetAllAccessKeys()

	assert.NoError(t, err)
	assert.Equal(t, mockData, data)
	mockDb.AssertCalled(t, "GetAllAccessData")
}

func TestGetAllAccessKeys_Error(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	mockDb.On("GetAllAccessData").Return(map[string]models.AccessKey{}, errors.New("db error"))

	data, err := service.GetAllAccessKeys()

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	assert.Empty(t, data)
	mockDb.AssertCalled(t, "GetAllAccessData")
}

func TestDeleteAccessKeys_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DeleteAccessData", keyId).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(nil)

	err := service.DeleteAccessKeys(keyId)

	assert.NoError(t, err)
	mockDb.AssertCalled(t, "DeleteAccessData", keyId)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestDeleteAccessKeys_DeleteAccessDataError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DeleteAccessData", keyId).Return(errors.New("delete error"))

	err := service.DeleteAccessKeys(keyId)

	assert.Error(t, err)
	assert.Equal(t, "delete error", err.Error())
	mockDb.AssertCalled(t, "DeleteAccessData", keyId)
	mockStream.AssertNotCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestDeleteAccessKeys_PublishError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DeleteAccessData", keyId).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(errors.New("publish error"))

	err := service.DeleteAccessKeys(keyId)

	assert.Error(t, err)
	assert.Equal(t, "publish error", err.Error())
	mockDb.AssertCalled(t, "DeleteAccessData", keyId)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestGetDataByAccessKey_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"
	expectedData := models.AccessKey{
		UserId:    1,
		RateLimit: 100,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
		Enabled:   true,
	}

	mockDb.On("GetAccessData", keyId).Return(expectedData, true)

	data, err := service.GetDataByAccessKey(keyId)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
	mockDb.AssertCalled(t, "GetAccessData", keyId)
}

func TestGetDataByAccessKey_NotFound(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("GetAccessData", keyId).Return(models.AccessKey{}, false)

	data, err := service.GetDataByAccessKey(keyId)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	assert.Empty(t, data)
	mockDb.AssertCalled(t, "GetAccessData", keyId)
}

func TestGetDataByAccessKey_DisabledKey(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"
	disabledData := models.AccessKey{
		UserId:    1,
		RateLimit: 100,
		Expiry:    time.Now().Add(time.Hour * 2).Unix(),
		Enabled:   false,
	}

	mockDb.On("GetAccessData", keyId).Return(disabledData, true)

	data, err := service.GetDataByAccessKey(keyId)

	assert.Error(t, err)
	assert.Equal(t, "key disabled", err.Error())
	assert.Empty(t, data)
	mockDb.AssertCalled(t, "GetAccessData", keyId)
}

func TestDisableAccessKey_Success(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DisableAccessKey", keyId).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(nil)

	err := service.DisableAccessKey(keyId)

	assert.NoError(t, err)
	mockDb.AssertCalled(t, "DisableAccessKey", keyId)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestDisableAccessKey_DisableAccessKeyError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DisableAccessKey", keyId).Return(errors.New("disable error"))

	err := service.DisableAccessKey(keyId)

	assert.Error(t, err)
	assert.Equal(t, "disable error", err.Error())
	mockDb.AssertCalled(t, "DisableAccessKey", keyId)
	mockStream.AssertNotCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}

func TestDisableAccessKey_PublishError(t *testing.T) {
	mockDb := new(testutils.MockDatabase)
	mockStream := new(testutils.MockStream)
	service := AccessKeyServices{
		db:     mockDb,
		stream: mockStream,
	}

	keyId := "test-key-id"

	mockDb.On("DisableAccessKey", keyId).Return(nil)
	mockStream.On("Publish", utils.PUBLISH_TOPIC, mock.Anything).Return(errors.New("publish error"))

	err := service.DisableAccessKey(keyId)

	assert.Error(t, err)
	assert.Equal(t, "publish error", err.Error())
	mockDb.AssertCalled(t, "DisableAccessKey", keyId)
	mockStream.AssertCalled(t, "Publish", utils.PUBLISH_TOPIC, mock.Anything)
}
