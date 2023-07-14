package biz

import (
	"context"

	"log"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/model"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *model.Like) error
}

//type IncreaseItemStorage interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeItemBiz struct {
	store UserLikeItemStore
	//itemStore IncreaseItemStorage
	ps pubsub.PubSub
}

func NewUserLikeItemBiz(
	store UserLikeItemStore,
	//itemStore IncreaseItemStorage,
	ps pubsub.PubSub,
) *userLikeItemBiz {
	return &userLikeItemBiz{
		store: store,
		//itemStore: itemStore,
		ps: ps,
	}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikedItem, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	//
	//if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
