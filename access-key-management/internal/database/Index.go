package database

import (
	"access-key-management/internal/models"
	"access-key-management/pkg/utils"
	"strings"
	"sync"
)

type Database interface {
	init() error
	SaveAccessData(key string, data models.AccessKey) error
	DeleteAccessData(key string) error
	UpdateAccessData(key string, data models.UpdateAccessKeyRequest) (models.AccessKey, error)
	GetAccessData(key string) (models.AccessKey, bool)
	DisableAccessKey(key string) error
	GetAllAccessData() (map[string]models.AccessKey, error)
}

var (
	database Database
	once     sync.Once
)

func GetDb() Database {
	once.Do(func() {
		if strings.EqualFold(utils.DB_TYPE, "local") {
			database = &Cache{}
		} else {
			database = &MySql{}
		}
		database.init()
	})
	return database
}
