package database

import (
	"database/sql"
	"web3-tokeninfo/internal/models"
)

type MySql struct {
	mysqlDb *sql.DB
}

func (f *MySql) init() error {
	var err error
	f.mysqlDb, err = sql.Open("mysql", "")
	return err
}

func (f *MySql) Ping() error {
	return nil
}

func (f *MySql) SaveAccessData(key string, data models.AccessKey) error {
	return nil
}

func (f *MySql) GetAccessData(key string) (models.AccessKey, bool) {
	return models.AccessKey{}, false
}

func (f *MySql) GetAllAccessData() (map[string]models.AccessKey, error) {
	return nil, nil
}

func (f *MySql) DeleteAccessData(key string) error {
	return nil
}

func (f *MySql) UpdateAccessData(key string, data models.UpdateAccessKeyRequest) error {
	return nil
}

func (f *MySql) DisableAccessKey(key string) error {
	return nil
}
