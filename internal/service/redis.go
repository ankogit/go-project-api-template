package service

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
)

type ConfigRedis struct {
	Host     string
	Port     string
	Db       int
	Password string
}

type RedisService struct {
	*redis.Client
}

func NewRedisService(cfg ConfigRedis) *RedisService {
	redisService := &RedisService{
		redis.NewClient(&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
			DB:       0,
		}),
	}

	_, err := redisService.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalf("Error trying to ping redis: %v", err)
	}
	return redisService
}
