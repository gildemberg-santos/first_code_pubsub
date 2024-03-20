package main

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	var topic string
	var message string

	flag.StringVar(&topic, "topic", "", "topic to publish to")
	flag.StringVar(&message, "message", "", "message to publish")
	flag.Parse()

	fmt.Println("Publishing to", topic)
	fmt.Println("Message", message)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.RPush(topic, message).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensagem enviada com sucesso")
}
