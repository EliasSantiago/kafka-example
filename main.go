package main

import (
	"log"

	"github.com/EliasSantiago/kafka-example/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("olá", "readtest", producer)

	for {
		_ = 1
	}
}
