package subscriber

import (
	"context"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/storage"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
	"gorm.io/gorm"
)

func DecreaseLikeCountAfterUserUnlikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Decrease like count after user unlikes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasItemId)

			return storage.NewSQLStore(db).DecreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
