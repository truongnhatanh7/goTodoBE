package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

// handler -> biz [-> repo] -> storage

type CreateItemStorage interface {
  CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
  store CreateItemStorage
}

// Any Storage has to implement CreateItem
func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
  return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
  if err := data.Validate(); err != nil {
    return err
  }

  if err := biz.store.CreateItem(ctx, data); err != nil {
    return common.ErrCannotCreateEntity(model.EntityName, err)
  }

  return nil
}