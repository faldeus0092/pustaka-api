package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	bookResponse := convertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// fungsi untuk membuat handler book yang akan membuat service
func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	// bind json agar bisa diproses, kirim &bookRequest untuk diubah
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// for storing multiple errors
		errorMessages := []string{}

		var ve validator.ValidationErrors
		var jsErr *json.UnmarshalTypeError

		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else if errors.As(err, &jsErr) {
			errorMessages = append(errorMessages, "Invalid JSON")
		} else {
			fmt.Println(err)
			errorMessages = append(errorMessages, "Something went wrong")
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	bookResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	// bind json agar bisa diproses, kirim &bookRequest untuk diubah
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// for storing multiple errors
		errorMessages := []string{}

		var ve validator.ValidationErrors
		var jsErr *json.UnmarshalTypeError

		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else if errors.As(err, &jsErr) {
			errorMessages = append(errorMessages, "Invalid JSON")
		} else {
			fmt.Println(err)
			errorMessages = append(errorMessages, "Something went wrong")
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	bookResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	bookResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
	}
}
