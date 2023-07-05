package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

type GetItemStorage interface {
  GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBiz struct {
  store GetItemStorage
}

// Any stores that are passed in have to implement GetItem - rule set by GetItemStorage interface
func NewGetItemBiz(store GetItemStorage) *getItemBiz {
  return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
  data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

  if err != nil {
    return nil, common.ErrCannotGetEntity(model.EntityName, err)
  }
  return data, nil
}