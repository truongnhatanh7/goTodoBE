package storage

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
