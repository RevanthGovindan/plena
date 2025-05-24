package database

import (
	"strings"
	"sync"
	"web3-tokeninfo/internal/models"
	"web3-tokeninfo/pkg/utils"
)

type Database interface {
	init() error
	Ping() error
	SaveAccessData(key string, data models.AccessKey) error
	DeleteAccessData(key string) error
	UpdateAccessData(key string, data models.UpdateAccessKeyRequest) error
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
