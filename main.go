package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
  Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

func main() {

	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)

	now := time.Now().UTC()
	item := TodoItem{
		Id:          1,
		Title:       "Task 1",
		Description: "Content 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsData, err := json.Marshal(item)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(jsData))

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id")
		}
	}

	r.Run(":3000")

}

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItemCreation

		if err := c.ShouldBind(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		res := db.Create(&itemData)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": itemData,
		})
	}
}

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItemCreation

    id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		res := db.Where("id = ?", id).First(&itemData)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": itemData,
		})
	}
}

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var updateData TodoItemUpdate

    id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

    if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		res := db.Where("id = ?", id).Updates(&updateData)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
