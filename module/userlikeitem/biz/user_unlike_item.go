package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/model"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStore
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStore) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store}
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

	return nil
}
