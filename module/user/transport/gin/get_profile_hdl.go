package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
