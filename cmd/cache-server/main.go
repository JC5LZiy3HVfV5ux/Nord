package main

import (
	"log"
	"os"
	"time"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/server"
)

func init() {
	file, err := os.OpenFile("/var/log/server/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Println("Не создался файл логов: ", err)
		return
	}

	log.SetOutput(file)

	log.Println("Создание логов... ok")
}

func main() {
	expiration, err := time.ParseDuration("10m")
	if err != nil {
		log.Fatal(err)
	}

	app, err := server.NewServer(&server.ServerConfig{
		Host:            "0.0.0.0",
		Port:            5000,
		Key:             "3bf9b8e0d3e815effba9f52b82559220",
		Lang:            "ru",
		Unit:            "metric",
		RedisHost:       "redis",
		RedisPort:       6379,
		RedisPassword:   "1234",
		RedisDB:         0,
		RedisExpiration: expiration,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := app.StartServer(); err != nil {
		log.Fatal(err)
	}
}
