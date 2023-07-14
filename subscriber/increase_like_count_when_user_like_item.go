package subscriber

import (
	"context"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/storage"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
	"gorm.io/gorm"
)

type HasItemId interface {
	GetItemId() int
}

//func IncreaseLikeCountAfterUserLikeItem(serviceCtx goservice.ServiceContext, ctx context.Context) {
//	ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)
//	db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
//
//	c, _ := ps.Subscribe(ctx, common.TopicUserLikedItem)
//
//	go func() {
//		defer common.Recovery()
//		for msg := range c {
//			data := msg.Data().(HasItemId)
//
//			if err := storage.NewSQLStore(db).IncreaseLikeCount(ctx, data.GetItemId()); err != nil {
//				log.Println(err)
//			}
//		}
//	}()
//}

func IncreaseLikeCountAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Increase like count after user likes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasItemId)

			return storage.NewSQLStore(db).IncreaseLikeCount(ctx, data.GetItemId())
		},
	}
}
