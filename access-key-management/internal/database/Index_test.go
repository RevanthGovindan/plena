package database

import (
	"access-key-management/internal/global"
	"sync"
	"testing"
)

func TestGetDb_LocalDatabase(t *testing.T) {
	// Set DbType to "local"
	global.Config.DbType = "local"

	// Call GetDb and assert the type of the returned database
	db := GetDb()
	if _, ok := db.(*Cache); !ok {
		t.Errorf("Expected *Cache, got %T", db)
	}
}

func TestGetDb_MySqlDatabase(t *testing.T) {
	// to reset again
	once = sync.Once{}
	// Set DbType to something other than "local"
	global.Config.DbType = "mysql"

	// Call GetDb and assert the type of the returned database
	db := GetDb()
	if _, ok := db.(*MySql); !ok {
		t.Errorf("Expected *MySql, got %T", db)
	}
}

func TestGetDb_Singleton(t *testing.T) {

	// Set DbType to "local"
	global.Config.DbType = "local"

	// Call GetDb multiple times and assert the same instance is returned
	db1 := GetDb()
	db2 := GetDb()
	if db1 != db2 {
		t.Errorf("Expected singleton instance, but got different instances")
	}
}
