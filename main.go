package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jaedung/go-fiber-tutorial/book"
	"github.com/jaedung/go-fiber-tutorial/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/books", book.NewBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	initDatabase()
	defer database.DBConn.Close()

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.SendStatus(200)
	// })

	// app.Get("/", helloWorld)

	setupRoutes(app)

	app.Listen(3000)
}
