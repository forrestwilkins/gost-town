package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	return db
}

func main() {
	app := fiber.New()
	db := initDB()

	app.Get("/read", func(c *fiber.Ctx) error {
		var user User

		err := db.Where(&User{Name: "nomad 1"}).Take(&user).Error
		if err != nil {
			c.SendString(err.Error())
		}

		return c.SendString(user.Name)
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		user := User{Name: "nomad 1"}
		result := db.Create(&user)

		fmt.Println(result.Error)
		return c.SendString(user.Name)
	})

	app.Listen(":3000")
}
