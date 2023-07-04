package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

// handler -> biz [-> repo] -> storage

type CreateItemStorage interface {
  CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
  store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
  return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
  if err := data.Validate(); err != nil {
    return err
  }

  if err := biz.store.CreateItem(ctx, data); err != nil {
    return err
  }

  return nil
}