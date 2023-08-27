package oauth

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type OAuthCache struct {
	client *redis.Client
	expire time.Duration
}

func (cache *OAuthCache)Init(url string,db int,expire time.Duration,password string){
	cache.client=redis.NewClient(&redis.Options{
        Addr:     url,
        Password: password, // no password set
        DB:       db,  // use default DB
    })
	cache.expire=expire
}

func (cache *OAuthCache)SetCache(userID,token string)(error){
	return cache.client.Set(cache.client.Context(), "token:"+token, userID, cache.expire).Err()
}

func (cache *OAuthCache)RemoveCache(token string){
  cache.client.Del(cache.client.Context(), "token:"+token)
}

func (cache *OAuthCache)GetUserID(token string)(string,error){
	return cache.client.Get(cache.client.Context(), "token:"+token).Result()
}

