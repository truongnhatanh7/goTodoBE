package subscriber

import (
	"context"

	"log"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
)

type HasUserId interface {
	GetUserId() int
}

func PushNotificationAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Title: "Push notification after user likes item",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			//db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

			data := message.Data().(HasUserId)

			log.Println("Push notification to user id:", data.GetUserId())

			return nil
		},
	}
}
