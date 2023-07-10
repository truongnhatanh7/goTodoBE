package storage

import (
	"context"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/module/userlikeitem/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, userId, itemId int) (*model.Like, error) {
	var data model.Like

	if err := s.db.Where("user_id = ? and item_id = ?", userId, itemId).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}