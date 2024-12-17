package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// ตั้งค่า Gin และ CORS
	r := gin.Default()
	r.Use(cors.Default()) // ใช้ CORS เริ่มต้นเพื่ออนุญาตให้ทุกแหล่งที่มาสามารถเข้าถึงได้

	// เชื่อมต่อกับฐานข้อมูล SQLite
	db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// สร้างหรืออัปเดตตาราง Books ในฐานข้อมูล
	db.AutoMigrate(&Book{})

	// สร้าง routes สำหรับ API
	r.GET("/books", getBooks)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	r.Run(":3000")
}

// โครงสร้างข้อมูล Book
type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ฟังก์ชันเพื่อดึงข้อมูลหนังสือทั้งหมด
func getBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// ฟังก์ชันเพื่อเพิ่มหนังสือใหม่
func createBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// ฟังก์ชันเพื่ออัปเดตข้อมูลหนังสือ
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// ฟังก์ชันเพื่อลบข้อมูลหนังสือ
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Book{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
