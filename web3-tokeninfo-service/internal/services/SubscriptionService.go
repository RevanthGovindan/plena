package services

import (
	"encoding/json"
	"fmt"
	"web3-tokeninfo/internal/database"
	"web3-tokeninfo/internal/models"
	"web3-tokeninfo/pkg/utils"
)

func HandleEvents(msg string) {
	var payload models.EventMessage
	err := json.Unmarshal([]byte(msg), &payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch payload.Event {
	case utils.ACCESSKEY_CREATED:
		handleCreation(payload)
	case utils.ACCESSKEY_DELETED:
		handleDeletion(payload)
	case utils.ACCESSKEY_UPDATED:
		handleUpdate(payload)
	case utils.ACCESSKEY_DISABLED:
		handleDisable(payload)
	}
}

func handleCreation(payload models.EventMessage) error {
	payload.Data.Enabled = true
	return database.GetDb().SaveAccessData(payload.Data.KeyId, payload.Data)
}

func handleDeletion(payload models.EventMessage) error {
	return database.GetDb().DeleteAccessData(payload.Data.KeyId)
}

func handleUpdate(payload models.EventMessage) error {
	database.LimiterStore.UpdateRateLimiter(payload.Data.KeyId, payload.Data.RateLimit)
	return database.GetDb().UpdateAccessData(payload.Data.KeyId, payload.Data)
}

func handleDisable(payload models.EventMessage) error {
	return database.GetDb().DisableAccessKey(payload.Data.KeyId)
}
