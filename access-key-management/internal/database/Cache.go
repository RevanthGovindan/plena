package database

import (
	"access-key-management/internal/models"
	"sync"
)

type Cache struct {
	mu         *sync.RWMutex
	accessData map[string]models.AccessKey
}

var (
	cache *Cache
	once  sync.Once
)

func GetDb() *Cache {
	once.Do(func() {
		cache = &Cache{accessData: make(map[string]models.AccessKey), mu: &sync.RWMutex{}}
	})
	return cache
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
