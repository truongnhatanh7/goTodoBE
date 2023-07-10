package ginuserlikeitem

import (
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/biz"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/storage"
	"gorm.io/gorm"
)

func ListUserLiked(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var queryString struct {
			common.Paging
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		queryString.Process()

		//requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		store := storage.NewSQLStore(db)
		business := biz.NewListUserLikeItemBiz(store)

		result, err := business.ListUserLikedItem(c.Request.Context(), int(id.GetLocalID()), &queryString.Paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, nil))
	}
}
