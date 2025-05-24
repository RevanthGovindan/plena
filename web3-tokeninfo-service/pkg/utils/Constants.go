package utils

const (
	STREAM_TYPE     = "redis"
	SUBSCRIBE_TOPIC = "events"
	DB_TYPE         = "local"
)

// events
const (
	ACCESSKEY_CREATED  = "access_key_created"
	ACCESSKEY_DELETED  = "access_key_deleted"
	ACCESSKEY_UPDATED  = "access_key_updated"
	ACCESSKEY_DISABLED = "access_key_disabled"
)

const (
	Response = `{ "data": { "id": "eth_0xdac17f958d2ee523a2206206994597c13d831ec7", "type": "token", "attributes": { "address": "0xdac17f958d2ee523a2206206994597c13d831ec7", "name": "Tether USD", "symbol": "USDT", "image_url": "https://assets.coingecko.com/coins/images/325/small/Tether.png?1696501661", "coingecko_coin_id": "tether", "websites": [ "https://tether.to/" ], "description": "Tether (USDT) is a cryptocurrency with a value meant to mirror the value of the U.S. dollar...", "gt_score": 92.6605504587156, "gt_score_details": { "pool": 87.5, "transaction": 0, "creation": 100, "info": 100, "holders": 0 }, "discord_url": null, "telegram_handle": null, "twitter_handle": "Tether_to", "categories": [], "gt_categories_id": [], "holders": { "count": 7041203, "distribution_percentage": { "top_10": "45.5782", "11_30": "13.4293", "31_50": "3.9681", "rest": "37.0244" }, "last_updated": "2025-03-12T05:28:50Z" }, "mint_authority": null, "freeze_authority": null } } }`
)
