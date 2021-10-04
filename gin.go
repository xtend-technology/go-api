package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB
var err error

func main() {

	r := gin.Default()

	USER := "root"
	PASS := "secret"
	HOST := "mysql"
	DBNAME := "test"

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Product{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		log.Println("Connection Established")
		// Get the first record ordered by primary key

	}

	r.GET("/", GetProducts)

	r.Run()
}

func GetProducts(c *gin.Context) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
