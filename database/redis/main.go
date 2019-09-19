package main

import (
	"log"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",
		DB:       0,
	})
	{
		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err)
		}
		if pong != "PONG" {
			log.Fatal("redis not respond PONG")
		}
	}
	{
		err := client.Set("key", "value", 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		val, err := client.Get("key").Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("key", val)
	}

	{
		val2, err := client.Get("key2").Result()

		if err == redis.Nil {
			log.Println("key2 does not exist")

		} else if err != nil {
			log.Fatal(err)

		} else {
			log.Println("key2", val2)
		}

	}
}

