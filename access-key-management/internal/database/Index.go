package database

import (
	"access-key-management/internal/global"
	"access-key-management/internal/models"
	"strings"
	"sync"
)

type Database interface {
	Init() error
	Ping() error
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
		if strings.EqualFold(global.Config.DbType, "local") {
			database = &Cache{}
		} else {
			database = &MySql{}
		}
		database.Init()
	})
	return database
}
