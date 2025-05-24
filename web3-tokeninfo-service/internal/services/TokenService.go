package services

import (
	"encoding/json"
	"fmt"
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

func handleCreation(payload models.EventMessage) {
	fmt.Println("handleCreation", payload)
}

func handleDeletion(payload models.EventMessage) {
	fmt.Println("handleDeletion", payload)
}

func handleUpdate(payload models.EventMessage) {
	fmt.Println("handleUpdate", payload)
}

func handleDisable(payload models.EventMessage) {
	fmt.Println("handleDisable", payload)
}
