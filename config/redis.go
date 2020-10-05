package config

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {

	redisHost := fatalGetString("REDIS_HOST")
	redisPort := fatalGetString("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Error(err)
	}

	log.Println("Redis Client Connected %s", pong)

	return client

}
