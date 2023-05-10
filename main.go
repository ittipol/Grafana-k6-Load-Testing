package main

import (
	"fmt"
	"go-load-testing/database"
	"go-load-testing/repositories"
	"go-load-testing/services"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {

	initTimeZone()

	db := initDbConnection()
	redisClient := initRedisConnection()

	// productRepositoryDbRepo := repositories.NewProductRepositoryDB(db)
	// productMongoDbRepo := repositories.NewProductMongoDB(any)
	productRedisRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	productService := services.NewProductService(productRedisRepo)

	productService.GetProduct()

	// app := fiber.New()

	// app.Get("/health", func(c *fiber.Ctx) error {
	// 	// time.Sleep(time.Second * 1)
	// 	return c.JSON("OK")
	// })

	// app.Listen(":5000")

}

func initTimeZone() {
	// LoadLocation looks for the IANA Time Zone database
	// List of tz database time zones
	// https: //en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	// init system time zone
	time.Local = location

	// timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(timeInUTC.In(location))
}

func initDbConnection() *gorm.DB {
	return database.GetDbConnection(
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.db"),
		false,
	)
}

func initRedisConnection() *redis.Client {
	return database.GetRedisConnection(
		viper.GetString("redis.username"),
		viper.GetString("redis.password"),
		viper.GetString("redis.host"),
		viper.GetInt("redis.port"),
		viper.GetInt("redis.db"),
	)
}

func initMongoDbConnection() *mongo.Client {
	return database.GetMongoDbConnection(
		viper.GetString("redis.username"),
		viper.GetString("redis.password"),
		viper.GetString("redis.host"),
		viper.GetInt("redis.port"),
	)
}
