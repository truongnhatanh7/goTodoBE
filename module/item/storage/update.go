package storage

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
  if err := s.db.Where(cond).Updates(&data).Error; err != nil {
    return err
  }

  return nil
}