package main

import (
	"database/sql"
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	_ "github.com/sijms/go-ora/v2"
)

func main() {
	// Oracle DB Connection
	conn, err := sql.Open("oracle", "oracle://fleet:teelf@192.168.1.77:1521/world")
	// check for error
	// if err != nil {
	// 	log.Fatal("DB Connection Error: ", err)
	// }

	// fmt.Println("DB Connection Successfully established")
	defer conn.Close()

	// stmt, err := conn.Prepare("SELECT col_1, col_2, col_3 FROM table WHERE col_1 = :1 or col_2 = :2")
	// // check for error
	// defer stmt.Close()

	// MySQL DB Connection
	// dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	// _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB Connection Error: ", err)
	// }

	// fmt.Println("DB Connection Successfully established")

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run("localhost:8088")
}
