package utils

import (
	"encoding/json"
	"fire_heart/models/msgraph"
	"github.com/patrickmn/go-cache"
	"time"
)

var GlobalCache *cache.Cache
var Ms365Token string

func InitCache() {
	GlobalCache = cache.New(50*time.Minute, 100*time.Minute)
}

func GetToken()  {
	authenticatedCache, found := GlobalCache.Get("authenticated")
	if found {
		authenticated := &msgraph.Token{}
		_ = json.Unmarshal([]byte(authenticatedCache.(string)), authenticated)
		Ms365Token = authenticated.AccessToken
	} else {
		Ms365Token = ""
	}
}
