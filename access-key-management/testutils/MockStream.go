package testutils

import (
	"access-key-management/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockStream struct {
	mock.Mock
}

func (m *MockStream) Init() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockStream) Publish(topic string, message models.EventMessage) error {
	args := m.Called(topic, message)
	return args.Error(0)
}

func (m *MockStream) Ping() error {
	args := m.Called()
	return args.Error(0)
}
