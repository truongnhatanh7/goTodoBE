package ginuserlikeitem

import (
	"net/http"
	"time"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/biz"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/model"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/storage"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
	"gorm.io/gorm"
)

func LikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

		store := storage.NewSQLStore(db)
		//itemStore := itemStorage.NewSQLStore(db)
		business := biz.NewUserLikeItemBiz(store, ps)
		now := time.Now().UTC()

		if err := business.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    int(id.GetLocalID()),
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
