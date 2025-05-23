package services

import (
	"access-key-management/internal/models"
	"access-key-management/pkg/utils"
	"time"
)

func CreateNewAccessKeys() (models.AccessKeyResponse, error) {
	expTime := time.Now().Add(time.Hour * 2).Unix()
	var response = models.AccessKeyResponse{
		AccessKey: models.AccessKey{UserId: utils.GenerateRandom(),
			RateLimit: 100,
			Expiry:    expTime,
		},
	}
	token, err := utils.GenerateToken(map[string]interface{}{
		"userId":    response.UserId,
		"rateLimit": response.RateLimit,
		"exp":       response.Expiry,
	})
	if err != nil {
		return models.AccessKeyResponse{}, err
	}
	response.Token = token
	return response, nil
}
