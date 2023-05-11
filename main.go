package main

import (
	"fmt"
	"go-load-testing/database"
	"go-load-testing/encryption"
	"go-load-testing/handlers"
	"go-load-testing/middlewares"
	"go-load-testing/repositories"
	"go-load-testing/services"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
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

	// Repository
	productRepo := repositories.NewProductRepositoryDB(db)

	// Service
	productService := services.NewProductServiceRedis(productRepo, redisClient)

	// Handler
	productHandler := handlers.NewProductHandler(productService)

	// ========================================================================================

	// privateKey, _ := encryption.GenerateRsaKeyPair()

	// encryption.Test()

	// ========================================================================================

	app := fiber.New()

	// Required. The key should be 32 bytes of random data in base64-encoded form.
	// You may run `openssl rand -base64 32` or use `encryptcookie.GenerateKey()` to generate a new key.
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "l/CsrINHL9WGXZdHMPRKesn8/jdblFYyOF8Ji0SW1lU=",
	}))

	v1 := app.Group("v1", middlewares.VersionV1, logger.New(logger.Config{
		Format: "[${time}] [${ip}]:${port} ${status} - ${latency} ${method} ${path}\n",
	}))

	v1.Post("/auth", func(c *fiber.Ctx) error {

		c.Cookie(&fiber.Cookie{
			Name:     "access-token",
			Value:    "Access Token",
			HTTPOnly: true,
			// Expires:  time.Now().Add(time.Second * 30),
			SessionOnly: true,
		})

		c.Cookie(&fiber.Cookie{
			Name:     "refresh-token",
			Value:    "Refresh Token",
			HTTPOnly: true,
			// Expires:  time.Now().Add(time.Minute * 30),
			SessionOnly: true,
		})

		c.Cookie(&fiber.Cookie{
			Name:     "public-key",
			Value:    "Public Key",
			HTTPOnly: true,
			// Expires:  time.Now().Add(time.Minute * 30),
			SessionOnly: true,
		})
		return c.JSON("OK")
	})

	v1.Post("/get-auth-profile", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return c.SendString("value=" + c.Cookies("access-token"))
	})

	productGroup := v1.Group("product", middlewares.Logger)

	productGroup.Use(middlewares.Auth).Get("/getProducts", productHandler.GetProducts)

	productGroup.Get("getBy/:id/:name?", func(c *fiber.Ctx) error {
		// Variable is now immutable
		result := utils.CopyString(c.Params("id"))

		c.SendStatus(fiber.StatusOK)
		return c.SendString(result)
	})

	app.Get("encryption", func(c *fiber.Ctx) error {

		_, publicKey := encryption.GenRsaKey()

		encryptedBytes := encryption.Encryption("Hello World 1234", publicKey)

		type res struct {
			Res    string
			Byte   []byte
			Base64 string
		}

		// Convert the encrypted bytes to encrypted text.
		dst := encryption.EncodeBase64(encryptedBytes)

		return c.JSON(res{
			Res:    string(encryptedBytes),
			Byte:   encryptedBytes,
			Base64: string(dst),
		})
	})

	// app.Get("/health", func(c *fiber.Ctx) error {
	// 	// time.Sleep(time.Second * 1)
	// 	return c.JSON("OK")
	// })

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

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
