package biz

import (
	"context"

	"log"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/model"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

//type DecreaseItemStorage interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStore
	//itemStore DecreaseItemStorage
	ps pubsub.PubSub
}

func NewUserUnlikeItemBiz(
	store UserUnlikeItemStore,
	//itemStore DecreaseItemStorage,
	ps pubsub.PubSub,
) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{
		store: store,
		//itemStore: itemStore,
		ps: ps,
	}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)

	// Delete if data existed
	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserUnlikedItem, pubsub.NewMessage(&model.Like{UserId: userId, ItemId: itemId})); err != nil {
		log.Println(err)
	}

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	if err := biz.itemStore.DecreaseLikeCount(ctx, itemId); err != nil {
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
