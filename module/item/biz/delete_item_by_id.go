package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store: store}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id":id}); err != nil {
		return err
	}

	return nil
}
