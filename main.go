package main

import (
	"fmt"

	"log"
	"os"

	"github.com/ElectronSz/bookapp/book"
	"github.com/ElectronSz/bookapp/database"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//set up fiber router routes
func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book", book.EditBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

//initialize database
func initDatabase() {
	var err error
	er := godotenv.Load()

	if er != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to database")
	}

	fmt.Println("Database connection initiated successfully")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	//new fiber app
	app := fiber.New()

	//init database
	initDatabase()

	//setup routes
	setUpRoutes(app)

	//start server and listen port 3000
	app.Listen(":3000")
}
