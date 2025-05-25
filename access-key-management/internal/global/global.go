package global

import "access-key-management/internal/models"

var (
	Config models.Config = models.Config{
		DbType:     "local",
		StreamType: "redis",
	}
)
