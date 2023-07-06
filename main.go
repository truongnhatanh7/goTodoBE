package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/truongnhatanh7/goTodoBE/component/uploadprovider"
	"github.com/truongnhatanh7/goTodoBE/middleware"
	ginitem "github.com/truongnhatanh7/goTodoBE/module/item/transport/gin"
	"github.com/truongnhatanh7/goTodoBE/module/upload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	bucketName := os.Getenv("G09BucketName")
	region := os.Getenv("G09Region")
	apiKey := os.Getenv("G09AccessKey")
	secret := os.Getenv("G09SecretKey")
	domain := ""

	s3Provider := uploadprovider.NewS3Provider(bucketName, region, apiKey, secret, domain)

	if err != nil {
		log.Fatalln(err)
	}

	// log.Println(db)

	r := gin.Default()
	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db, s3Provider))

		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.Run(":3000")

}
