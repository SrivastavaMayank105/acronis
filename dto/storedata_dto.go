package dto

import "time"

type StoredDataInfo struct {
	Key            string      `json:"key"`
	DataValue      interface{} `json:"datavalue"`
	CreationTime   time.Time   `json:"creation_time"`
	ExpirationTime time.Time   `json:"expiration_time"`
}
