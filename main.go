package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func initDB() {
	var err error
	// ใช้ GORM กับไดรเวอร์ SQLite ที่เหมาะสม
	db, err = gorm.Open(sqlite.Open("data/books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Book{})
}

func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func getBooks(c *gin.Context) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func main() {
	initDB()
	r := gin.Default()

	r.POST("/books", createBook)
	r.GET("/books", getBooks)

	if err := r.Run(":3000"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
