package models

type AccessKey struct {
	KeyId     string `json:"keyId"`
	UserId    int    `json:"userId"`
	RateLimit int    `json:"rateLimit"`
	Expiry    int64  `json:"expiry"`
	Enabled   bool   `json:"-"`
}

type EventMessage struct {
	Event string    `json:"event"`
	Data  AccessKey `json:"data"`
}
