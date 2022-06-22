package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	db.AutoMigrate(book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// CREATE
	// CreatedBook := book.Book{
	// 	Title:       "DARK SOULS 3",
	// 	Price:       599000,
	// 	Rating:      5,
	// 	Description: "One of the Best action RPG of All Time",
	// }

	// bookRepository.Create(CreatedBook)
	// bookRequest := book.BookRequest{
	// 	Title:       "DARK SOULS 2: Scholar of the First Sin",
	// 	Price:       "575000",
	// 	Rating:      "5",
	// 	Description: "Prequel of One of the Best action RPG of All Time",
	// }
	// bookService.Create(bookRequest)

	// if err != nil {
	// 	fmt.Println("Error creating book")
	// }

	// books, err := bookRepository.FindAll()
	// if err != nil {
	// 	fmt.Println("Error getting books: ", err)
	// }
	// for _, b := range books {
	// 	fmt.Println("Title: ", b.Title)
	// }

	// READ
	// variable to retrieve the book stuct
	// var book book.Book

	// Get the first data
	// err = db.Debug().First(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }
	// fmt.Println("Title: ", book.Title)
	// fmt.Println("book object: ", book)

	// Get specified/all data
	// var books []book.Book
	// // err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("title = ?", "ELDEN RING").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }
	// for _, v := range books {
	// 	fmt.Println("Title: ", v.Title)
	// 	fmt.Println("book object: ", v)
	// }

	// Update data
	// var book book.Book
	// // err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }
	// book.Title = "SEKIRO: Shadows Dies Twice"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }

	// Delete book
	// var book book.Book
	// // err = db.Debug().Find(&books).Error
	// err = db.Debug().Where("id = ?", 2).First(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error retrieving book")
	// 	fmt.Println("================================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", bookHandler.GetBook)
	v1.GET("/books", bookHandler.GetAllBooks)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
