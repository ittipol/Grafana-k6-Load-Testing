package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetRedisConnection(username string, password string, hostName string, port int, db int) *redis.Client {
	// redis://<user>:<pass>@localhost:6379/<db>

	dsn := fmt.Sprintf("redis://%v:%v@%v:%v/%v", username, password, hostName, port, db)
	println(dsn)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
