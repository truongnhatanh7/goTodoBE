package ginitem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/biz"
	"github.com/truongnhatanh7/goTodoBE/module/item/storage"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

    store := storage.NewSQLStore(db)
    business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)
    if err != nil {
      c.JSON(http.StatusBadRequest, err)

      return
    }

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
