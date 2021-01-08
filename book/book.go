package book

import (
	"github.com/ElectronSz/bookapp/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

//Book structure
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

//Response structure
type Response struct {
	Status  int
	Message string
}

//GetBooks : get all books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)

	c.JSON(books)
}

//GetBook : get single book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)

	c.JSON(book)
}

//NewBook : create a new Book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

//EditBook : edit one book
func EditBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(400).SendString(err.Error())
	}

	isupdated := db.Model(book).Updates(Book{Title: book.Title, Author: book.Author, Rating: book.Rating})

	if isupdated == nil {
		res := Response{Status: 400, Message: "Error Updating book"}
		c.Status(200).JSON(res)
	}
	res := Response{Status: 200, Message: "Book Edited Successfully"}
	c.Status(200).JSON(res)

}

//DeleteBook : delete one book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book

	db.First(&book, id)
	if book.Title == "" {
		res := Response{Status: 500, Message: "No book with given id found"}
		c.Status(500).JSON(res)
		return
	}

	db.Delete(&book)

	res := Response{Status: 200, Message: "Book deleted successfully"}
	c.JSON(res)
}
