package main

import (
	"time"
)

func main() {

	initTimeZone()
	// any := 1

	// Mysql
	// dsn := "root:1234@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db := database.GetDbConnection(dsn, false)

	// productRepositoryDbRepo := repositories.NewProductRepositoryDB(db)
	// productMongoDbRepo := repositories.NewProductMongoDB(any)
	// productRedisRepo := repositories.NewProductRedis(any)

	// services.NewProductService(productMongoDbRepo)

	// _ = productRepositoryDbRepo

	// products, _ := productRepositoryDbRepo.GetProduct()

	// fmt.Println(products)

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
