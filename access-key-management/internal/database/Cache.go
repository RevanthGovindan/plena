package database

import (
	"access-key-management/internal/models"
	"errors"
	"sync"
)

type Cache struct {
	mu         *sync.RWMutex
	accessData map[string]models.AccessKey
}

func (f *Cache) init() error {
	f.mu = &sync.RWMutex{}
	f.accessData = make(map[string]models.AccessKey)
	return nil
}

func (f *Cache) SaveAccessData(key string, data models.AccessKey) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.accessData[key] = data
	return nil
}

func (f *Cache) GetAccessData(key string) (models.AccessKey, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	keyData, exists := f.accessData[key]
	return keyData, exists
}

func (f *Cache) GetAllAccessData() (map[string]models.AccessKey, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.accessData, nil
}

func (f *Cache) DeleteAccessData(key string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.accessData, key)
	return nil
}

func (f *Cache) UpdateAccessData(key string, data models.UpdateAccessKeyRequest) (models.AccessKey, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	prevData, exists := f.accessData[key]
	if !exists {
		return models.AccessKey{}, errors.New("not found")
	}
	prevData.Expiry = data.Expiry
	prevData.RateLimit = data.RateLimit
	f.accessData[key] = prevData
	return prevData, nil
}

func (f *Cache) DisableAccessKey(key string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	prevData, exists := f.accessData[key]
	if !exists {
		return errors.New("not found")
	}
	if !prevData.Enabled {
		return errors.New("disabled already")
	}
	prevData.Enabled = false
	f.accessData[key] = prevData
	return nil
}
