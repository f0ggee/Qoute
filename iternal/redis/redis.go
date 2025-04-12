package redis

import (
	"context"
	_ "errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	_ "log"
	_ "os"
)

var ctx = context.Background()

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Username string `yaml:"username"`
}

func connect(ctx context.Context) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "AdA7AAIjcDEwOGYzODMxNjI0MDM0MTk3YmNjYzA2NDY0YWIwM2FiOHAxMA",
		Username: "Default",
	})
	if err := db.Ping(ctx).Err(); err != nil {
		fmt.Println(err)
		return nil, err

	}
	return db, nil

}
