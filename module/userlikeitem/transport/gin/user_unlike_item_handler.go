package ginuserlikeitem

import (
	"log"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	itemStorage "github.com/truongnhatanh7/goTodoBE/module/item/storage"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/biz"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/storage"
	"gorm.io/gorm"
)

func UnlikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		itemStore := itemStorage.NewSQLStore(db)
		business := biz.NewUserUnlikeItemBiz(store, itemStore)

		go func() {
			defer common.Recovery()

			if err := business.UnlikeItem(c.Request.Context(), requester.GetUserId(), int(id.GetLocalID())); err != nil {
				log.Println(err)
			}
		}()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
