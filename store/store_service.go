package store

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)


const CacheDuration = 24 * time.Hour

var (
	ctx = context.Background()
	storeService = StoreService{}
)

type StoreService struct {
	redisClient *redis.Client
}

func InitStoreService() *StoreService {
	redisClinet := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := redisClinet.Ping(ctx).Result()
	if err != nil {
		log.Panicf("Failed to connect to redis %v", err)
	}

	log.Printf("\n Redis Started Successfuly: pong message = {%s}" , pong)
	storeService.redisClient = redisClinet
	return &storeService
}

func SaveUrlMapping(shortUrl, longUrl, userId string)  {
	if err := storeService.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err(); err != nil {
		log.Printf("Failed to save url mapping %v", err)
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result , err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Printf("Failed to retrieve initial url: %v", err)
	}

	return result
}