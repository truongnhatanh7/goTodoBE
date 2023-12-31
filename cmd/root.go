package cmd

import (
	"fmt"
	"os"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/component/tokenprovider/jwt"
	"github.com/truongnhatanh7/goTodoBE/component/uploadprovider"
	"github.com/truongnhatanh7/goTodoBE/middleware"
	ginitem "github.com/truongnhatanh7/goTodoBE/module/item/transport/gin"
	"github.com/truongnhatanh7/goTodoBE/module/upload"
	userstorage "github.com/truongnhatanh7/goTodoBE/module/user/storage"
	ginuser "github.com/truongnhatanh7/goTodoBE/module/user/transport/gin"
	ginuserlikeitem "github.com/truongnhatanh7/goTodoBE/module/userlikeitem/transport/gin"
	"github.com/truongnhatanh7/goTodoBE/plugin/sdkgorm"
	"github.com/truongnhatanh7/goTodoBE/plugin/sdks3"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
	"github.com/truongnhatanh7/goTodoBE/subscriber"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.PluginDBMain)),
		goservice.WithInitRunnable(pubsub.NewPubSub(common.PluginPubSub)),
		goservice.WithInitRunnable(sdks3.NewS3Service("s3-provider", "aws-s3")),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {
		systemSecret := os.Getenv("SECRET")

		service := newService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			db := service.MustGet(common.PluginDBMain).(*gorm.DB)

			s3Provider := service.MustGet("aws-s3").(interface {
				GetBucketProvider() *uploadprovider.S3Provider
			}).GetBucketProvider()
			authStore := userstorage.NewSQLStore(db)
			tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
			middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

			v1 := engine.Group("/v1")
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

					items.POST("/:id/like", ginuserlikeitem.LikeItem(service))
					items.DELETE("/:id/unlike", ginuserlikeitem.UnlikeItem(service))
					items.GET("/:id/liked-users", ginuserlikeitem.ListUserLiked(service))
				}
			}
		})

		_ = subscriber.NewEngine(service).Start()

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
