package models

type AccessKey struct {
	UserId    int   `json:"userId"`
	RateLimit int   `json:"rateLimit"`
	Expiry    int64 `json:"expiry"`
}

type AccessKeyResponse struct {
	AccessKey
	Token string `json:"token"`
}
