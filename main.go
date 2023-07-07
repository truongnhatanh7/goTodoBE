package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/truongnhatanh7/goTodoBE/component/tokenprovider/jwt"
	"github.com/truongnhatanh7/goTodoBE/component/uploadprovider"
	"github.com/truongnhatanh7/goTodoBE/middleware"
	ginitem "github.com/truongnhatanh7/goTodoBE/module/item/transport/gin"
	"github.com/truongnhatanh7/goTodoBE/module/upload"
	"github.com/truongnhatanh7/goTodoBE/module/user/storage"
	ginuser "github.com/truongnhatanh7/goTodoBE/module/user/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	systemSecret := os.Getenv("SECRET")
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

	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	r := gin.Default()
	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db, s3Provider))

		v1.POST("/register", ginuser.Register(db))
		v1.POST("/login", ginuser.Login(db, tokenProvider))
		v1.GET("/profile", middlewareAuth, ginuser.Profile())

		items := v1.Group("/items", middlewareAuth)
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
