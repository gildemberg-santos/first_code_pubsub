package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	var topic string

	flag.StringVar(&topic, "topic", "", "topic to publish to")
	flag.Parse()

	fmt.Println("Subscribing to", topic)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	go func() {
		for {
			msg, err := client.BLPop(0*time.Second, topic).Result()
			if err != nil {
				panic(err)
			}

			fmt.Println("Received", msg[1])
		}
	}()

	select {}
}
