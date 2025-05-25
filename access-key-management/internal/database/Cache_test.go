package database

import (
	"access-key-management/internal/models"
	"testing"
)

func TestSaveAccessData(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}
	key := "test-key"
	data := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}

	err = cache.SaveAccessData(key, data)
	if err != nil {
		t.Fatalf("SaveAccessData returned an error: %v", err)
	}

	retrievedData, exists := cache.GetAccessData(key)
	if !exists {
		t.Fatalf("expected data to exist for key %s, but it does not", key)
	}

	if retrievedData != data {
		t.Errorf("retrieved data does not match saved data. Got %+v, want %+v", retrievedData, data)
	}
}

func TestGetAllAccessData(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}

	data1 := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}
	data2 := models.AccessKey{
		Expiry:    9876543210,
		RateLimit: 200,
		Enabled:   false,
	}

	err = cache.SaveAccessData("key1", data1)
	if err != nil {
		t.Fatalf("failed to save data for key1: %v", err)
	}

	err = cache.SaveAccessData("key2", data2)
	if err != nil {
		t.Fatalf("failed to save data for key2: %v", err)
	}

	allData, err := cache.GetAllAccessData()
	if err != nil {
		t.Fatalf("GetAllAccessData returned an error: %v", err)
	}

	if len(allData) != 2 {
		t.Fatalf("expected 2 items in cache, got %d", len(allData))
	}

	if allData["key1"] != data1 {
		t.Errorf("data for key1 does not match. Got %+v, want %+v", allData["key1"], data1)
	}

	if allData["key2"] != data2 {
		t.Errorf("data for key2 does not match. Got %+v, want %+v", allData["key2"], data2)
	}
}

func TestDisableAccessKey(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}

	key := "test-key"
	data := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}

	// Save initial data
	err = cache.SaveAccessData(key, data)
	if err != nil {
		t.Fatalf("failed to save access data: %v", err)
	}

	// Disable the access key
	err = cache.DisableAccessKey(key)
	if err != nil {
		t.Fatalf("DisableAccessKey returned an error: %v", err)
	}

	// Verify the key is disabled
	updatedData, exists := cache.GetAccessData(key)
	if !exists {
		t.Fatalf("expected data to exist for key %s, but it does not", key)
	}

	if updatedData.Enabled {
		t.Errorf("expected key %s to be disabled, but it is still enabled", key)
	}

	// Attempt to disable an already disabled key
	err = cache.DisableAccessKey(key)
	if err == nil || err.Error() != "disabled already" {
		t.Errorf("expected error 'disabled already', got %v", err)
	}

	// Attempt to disable a non-existent key
	err = cache.DisableAccessKey("non-existent-key")
	if err == nil || err.Error() != "not found" {
		t.Errorf("expected error 'not found', got %v", err)
	}
}

func TestUpdateAccessData(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}

	key := "test-key"
	initialData := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}

	// Save initial data
	err = cache.SaveAccessData(key, initialData)
	if err != nil {
		t.Fatalf("failed to save access data: %v", err)
	}

	// Update the access data
	updateRequest := models.UpdateAccessKeyRequest{
		Expiry:    9876543210,
		RateLimit: 200,
	}
	updatedData, err := cache.UpdateAccessData(key, updateRequest)
	if err != nil {
		t.Fatalf("UpdateAccessData returned an error: %v", err)
	}

	// Verify the updated data
	if updatedData.Expiry != updateRequest.Expiry || updatedData.RateLimit != updateRequest.RateLimit {
		t.Errorf("updated data does not match. Got %+v, want %+v", updatedData, updateRequest)
	}

	// Verify the data in the cache
	retrievedData, exists := cache.GetAccessData(key)
	if !exists {
		t.Fatalf("expected data to exist for key %s, but it does not", key)
	}

	if retrievedData.Expiry != updateRequest.Expiry || retrievedData.RateLimit != updateRequest.RateLimit {
		t.Errorf("retrieved data does not match updated data. Got %+v, want %+v", retrievedData, updateRequest)
	}

	// Attempt to update a non-existent key
	_, err = cache.UpdateAccessData("non-existent-key", updateRequest)
	if err == nil || err.Error() != "not found" {
		t.Errorf("expected error 'not found', got %v", err)
	}
}

func TestDeleteAccessData(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}

	key := "test-key"
	data := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}

	// Save initial data
	err = cache.SaveAccessData(key, data)
	if err != nil {
		t.Fatalf("failed to save access data: %v", err)
	}

	// Delete the access data
	err = cache.DeleteAccessData(key)
	if err != nil {
		t.Fatalf("DeleteAccessData returned an error: %v", err)
	}

	// Verify the data is deleted
	_, exists := cache.GetAccessData(key)
	if exists {
		t.Errorf("expected data for key %s to be deleted, but it still exists", key)
	}

	// Attempt to delete a non-existent key
	err = cache.DeleteAccessData("non-existent-key")
	if err != nil {
		t.Errorf("expected no error when deleting a non-existent key, got %v", err)
	}
}

func TestGetAccessData(t *testing.T) {
	cache := &Cache{}
	err := cache.Init()
	if err != nil {
		t.Fatalf("failed to initialize cache: %v", err)
	}

	key := "test-key"
	data := models.AccessKey{
		Expiry:    1234567890,
		RateLimit: 100,
		Enabled:   true,
	}

	// Save data to the cache
	err = cache.SaveAccessData(key, data)
	if err != nil {
		t.Fatalf("failed to save access data: %v", err)
	}

	// Retrieve the data
	retrievedData, exists := cache.GetAccessData(key)
	if !exists {
		t.Fatalf("expected data to exist for key %s, but it does not", key)
	}

	// Verify the retrieved data matches the saved data
	if retrievedData != data {
		t.Errorf("retrieved data does not match saved data. Got %+v, want %+v", retrievedData, data)
	}

	// Attempt to retrieve data for a non-existent key
	_, exists = cache.GetAccessData("non-existent-key")
	if exists {
		t.Errorf("expected no data to exist for non-existent key, but data was found")
	}
}
