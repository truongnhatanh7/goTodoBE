package storage

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
  var data model.TodoItem

  if err := s.db.Where(cond).First(&data).Error; err != nil {
    return nil, err
  }

  return &data, nil
}