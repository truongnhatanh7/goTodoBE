package storage

import (
	"context"

	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "Deleted"

	if err := s.db.
		Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
