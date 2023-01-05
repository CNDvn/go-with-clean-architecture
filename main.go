package main

import (
	"log"
	"os"
	"time"

	todotrpt "go-with-clean-architecture/module/item/transport"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

func main() {
	// load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalln("load env file error: ", err)
	}
	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	if !ok {
		log.Fatalln("Missing MySQL connection string")
	}

	dsn := mysqlConnStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected:", db)
	db.AutoMigrate(&ToDoItem{})

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", todotrpt.HandleCreateItem(db)) // create item
		// v1.GET("/items", getListOfItems(db))        // list items
		// v1.GET("/items/:id", readItemById(db))      // get an item by ID
		// v1.PUT("/items/:id", editItemById(db))      // edit an item by ID
		// v1.DELETE("/items/:id", deleteItemById(db)) // delete an item by ID
	}

	router.Run()
}

// func readItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataItem ToDoItem

// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": dataItem})
// 	}
// }

// func getListOfItems(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		type DataPaging struct {
// 			Page  int   `json:"page" form:"page"`
// 			Limit int   `json:"limit" form:"limit"`
// 			Total int64 `json:"total" form:"-"`
// 		}

// 		var paging DataPaging

// 		if err := c.ShouldBind(&paging); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if paging.Page <= 0 {
// 			paging.Page = 1
// 		}

// 		if paging.Limit <= 0 {
// 			paging.Limit = 10
// 		}

// 		offset := (paging.Page - 1) * paging.Limit

// 		var result []ToDoItem

// 		if err := db.Table(ToDoItem{}.TableName()).
// 			Count(&paging.Total).
// 			Offset(offset).
// 			Order("id desc").
// 			Find(&result).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": result})
// 	}
// }

// func editItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var dataItem ToDoItem

// 		if err := c.ShouldBind(&dataItem); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }

// func deleteItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Table(ToDoItem{}.TableName()).
// 			Where("id = ?", id).
// 			Delete(nil).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }
