package ginitem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/biz"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
	"github.com/truongnhatanh7/goTodoBE/module/item/storage"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// Paging
    var queryString struct {
      common.Paging
      model.Filter
    }

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		queryString.Paging.Process()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

    store := storage.NewSQLStore(db)
    business := biz.NewLIstItemBiz(store, requester)

    result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
        "error": err.Error(),
      })

      return
    }

		for i := range result {
			result[i].Mask()
		}


		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, queryString.Filter))
	}
}