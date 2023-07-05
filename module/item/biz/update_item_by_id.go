package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

type updateItem struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItem {
	return &updateItem{store: store}
}

func (biz *updateItem) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id":id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
