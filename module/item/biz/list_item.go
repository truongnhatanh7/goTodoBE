package biz

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

type ListItemStorage interface {
  ListItem(
    ctx context.Context,
    filter *model.Filter,
    paging *common.Paging,
    moreKeys ...string,
  ) ([]model.TodoItem, error)
}

type listItemBiz struct {
  store ListItemStorage
}

func NewLIstItemBiz(store ListItemStorage) *listItemBiz {
  return &listItemBiz{store: store}
}

func (biz*listItemBiz) ListItem(
  ctx context.Context,
  filter *model.Filter,
  paging *common.Paging,
) ([]model.TodoItem, error) {
  data, err := biz.store.ListItem(ctx, filter, paging)

  if err != nil {
    return nil, common.ErrCannotListEntity(model.EntityName, err)
  }

  return data, nil
}