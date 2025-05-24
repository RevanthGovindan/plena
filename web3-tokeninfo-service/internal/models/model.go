package models

type AccessKey struct {
	UserId    int   `json:"userId"`
	RateLimit int   `json:"rateLimit"`
	Expiry    int64 `json:"expiry"`
	Enabled   bool  `json:"-"`
}

type AccessKeyResponse struct {
	AccessKey
	KeyId string `json:"keyId"`
}

type EventMessage struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

type UpdateAccessKeyRequest struct {
	RateLimit int   `json:"rateLimit"`
	Expiry    int64 `json:"expiry"`
}
