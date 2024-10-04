package main

import (
	"hex/model"
	"hex/route"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=admin password=1234 dbname=golang_test port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("falied to connect database")
	}

	route.InitRoute(app, db)

	db.AutoMigrate(&model.Order{})

	app.Listen(":8000")
}
