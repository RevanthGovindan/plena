package testutils

import (
	"access-key-management/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of the Database interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Init() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) SaveAccessData(key string, data models.AccessKey) error {
	args := m.Called(key, data)
	return args.Error(0)
}

func (m *MockDatabase) DeleteAccessData(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockDatabase) UpdateAccessData(key string, data models.UpdateAccessKeyRequest) (models.AccessKey, error) {
	args := m.Called(key, data)
	return args.Get(0).(models.AccessKey), args.Error(1)
}

func (m *MockDatabase) GetAccessData(key string) (models.AccessKey, bool) {
	args := m.Called(key)
	return args.Get(0).(models.AccessKey), args.Bool(1)
}

func (m *MockDatabase) DisableAccessKey(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockDatabase) GetAllAccessData() (map[string]models.AccessKey, error) {
	args := m.Called()
	return args.Get(0).(map[string]models.AccessKey), args.Error(1)
}
