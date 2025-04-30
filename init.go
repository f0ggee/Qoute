package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load() вернёт nil, если .env найден,
	// и ошибку, если файла нет — сообщаем об этом, но не падаем.
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("godotenv: %v", err)
	}
}
