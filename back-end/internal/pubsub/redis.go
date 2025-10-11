package pubsub

import (
	"context"
	"fmt"
	"url-shortener/back-end/config"

	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisPubSub(cfg *config.Config) *RedisPubSub {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}

	return &RedisPubSub{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func (ps *RedisPubSub) Publish(channel string, message string) error {
	return ps.Client.Publish(ps.Ctx, channel, message).Err()
}

func (ps *RedisPubSub) Subscribe(channel string, handler func(msg string)) (func(), error) {
	pubsub := ps.Client.Subscribe(context.Background(), channel)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := pubsub.Channel()
		for {
			select {
			case msg := <-ch:
				handler(msg.Payload)
			case <-ctx.Done():
				return
			}
		}
	}()

	return func() {
		cancel()
		pubsub.Close()
	}, nil
}
