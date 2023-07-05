package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
	ginitem "github.com/truongnhatanh7/goTodoBE/module/item/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.Run(":3000")

}

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var updateData model.TodoItemUpdate

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

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		deletedStatus := "Deleted"

		res := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(&model.TodoItemUpdate{
			Status: &deletedStatus,
		})
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// Paging
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		//DB
		var result []model.TodoItem

		db = db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")

		if err := db.Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		res := db.
			Select("*").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result)

		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": res.Error.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
