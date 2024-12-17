# Go REST API - Book Management

This is a simple REST API for managing books, built with Go, Gin, and GORM, and exposes endpoints for performing CRUD (Create, Read, Update, Delete) operations on books.

## Requirements

- Go (version 1.18 or higher)
- Docker (optional for running the database)
- SQLite or MySQL database (configured in the code)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/go-rest-api.git
    cd go-rest-api
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the API:
    ```bash
    go run main.go
    ```

4. The server will start on `http://localhost:3000`.

## API Endpoints

### 1. Create a New Book
- **Endpoint:** `POST /books`
- **Description:** Creates a new book in the database.
- **Request Body:**
    ```json
    {
      "title": "Go Programming",
      "author": "John Doe"
    }
    ```
- **Example cURL Command:**
    ```bash
    curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title": "Go Programming", "author": "John Doe"}'
    ```
- **Response:**
    ```json
    {
      "id": 1,
      "title": "Go Programming",
      "author": "John Doe"
    }
    ```

### 2. Get All Books
- **Endpoint:** `GET /books`
- **Description:** Fetches all books from the database.
- **Example cURL Command:**
    ```bash
    curl http://localhost:3000/books
    ```
- **Response:**
    ```json
    [
      {
        "id": 1,
        "title": "Go Programming",
        "author": "John Doe"
      }
    ]
    ```

### 3. Get a Single Book by ID
- **Endpoint:** `GET /books/:id`
- **Description:** Fetches a single book by its ID.
- **Example cURL Command:**
    ```bash
    curl http://localhost:3000/books/1
    ```
- **Response:**
    ```json
    {
      "id": 1,
      "title": "Go Programming",
      "author": "John Doe"
    }
    ```

### 4. Update a Book by ID
- **Endpoint:** `PUT /books/:id`
- **Description:** Updates a book's details by its ID.
- **Request Body:**
    ```json
    {
      "title": "Go Programming Advanced",
      "author": "John Doe"
    }
    ```
- **Example cURL Command:**
    ```bash
    curl -X PUT http://localhost:3000/books/1 -H "Content-Type: application/json" -d '{"title": "Go Programming Advanced", "author": "John Doe"}'
    ```
- **Response:**
    ```json
    {
      "id": 1,
      "title": "Go Programming Advanced",
      "author": "John Doe"
    }
    ```

### 5. Delete a Book by ID
- **Endpoint:** `DELETE /books/:id`
- **Description:** Deletes a book by its ID.
- **Example cURL Command:**
    ```bash
    curl -X DELETE http://localhost:3000/books/1
    ```
- **Response:**
    ```json
    {
      "message": "Book deleted successfully"
    }
    ```

## Database Setup

If you're using SQLite, the database file will be created automatically when you run the API.

### Option 1: Run with SQLite
SQLite is used by default for this API. The database file will be created in the project directory (`books.db`).

### Option 2: Run with MySQL
If you prefer to use MySQL, change the database configuration in `main.go`:
```go
dsn := "user:password@tcp(localhost:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"



### สิ่งที่ประกอบใน `README.md` นี้:
1. **การติดตั้งและใช้งาน:**
   - ขั้นตอนในการติดตั้งโปรเจกต์
   - วิธีการรัน API
2. **API Endpoints:**
   - ตัวอย่างของแต่ละ endpoint เช่น `POST /books`, `GET /books`, `PUT /books/:id`, `DELETE /books/:id`
   - คำสั่ง `curl` สำหรับทดสอบ
3. **การตั้งค่าฐานข้อมูล:**
   - วิธีการตั้งค่าฐานข้อมูล (SQLite และ MySQL)
4. **การทดสอบ API:**
   - ตัวอย่างการทดสอบด้วย `curl` หรือ Postman
5. **Docker Setup:**
   - การใช้งาน Docker เพื่อรัน API และฐานข้อมูล

เมื่อคุณมีไฟล์ `README.md` ที่อธิบายการทำงานของ API นี้แล้ว ผู้ใช้คนอื่นๆ ที่เข้ามาในโปรเจกต์ของคุณจะสามารถเข้าใจและใช้งาน API ได้อย่างง่ายดายครับ!

