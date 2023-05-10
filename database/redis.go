package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetRedisConnection() {
	// redis://<user>:<pass>@localhost:6379/<db>

	dsn := fmt.Sprint("redis://%v:%v@%v:%v/%v")

	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
}
