package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
)

func Recover() gin.HandlerFunc {
  return func(c *gin.Context) {
    defer func() {
      if err := recover(); err != nil {
        c.Header("Content-Type", "application/json")

        // App error
        if appErr, ok := err.(*common.AppError); ok {
          c.AbortWithStatusJSON(appErr.StatusCode, appErr)
          // panic(err)
          return
        }

        // Primitive error
        appErr := common.ErrInternal(err.(error))
        c.AbortWithStatusJSON(appErr.StatusCode, appErr)
        // panic(err)
        return
      }  
    }()

    c.Next()
  }
}