package cache

import "todoapp/src/models"

type CacheData struct {
	OtpCode string
	NewUser models.UserCreate
}

var Cache = make(map[string]CacheData)
