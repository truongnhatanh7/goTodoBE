package model

import (
	"errors"
	"strings"

	"github.com/truongnhatanh7/goTodoBE/common"
)

var (
	ErrTitleCannotBeEmpty = errors.New("Title cannot be empty")
	ErrItemIsDeleted      = errors.New("Item is deleted")
)

const (
	EntityName = "item"
)

type TodoItem struct {
	common.SQLModel
	Title       string             `json:"title" gorm:"column:title;"`
	UserId      int                `json:"user_id" gorm:"column:user_id;"`
	Description string             `json:"description" gorm:"column:description;"`
	Status      string             `json:"status" gorm:"column:status;"`
	Image       *common.Image      `json:"image" gorm:"column:image;"`
	LikedCount  int                `json:"liked_count" gorm:"-"`
	Owner       *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"`
}

func (TodoItem) TableName() string { return "todo_items" }

func (i *TodoItem) Mask() {
	i.SQLModel.Mask(common.DbTypeItem)

	if v := i.Owner; v != nil {
		v.Mask()
	}
}

type TodoItemCreation struct {
	Id          int           `json:"id" gorm:"column:id;"`
	UserId      int           `json:"-" gorm:"column:user_id;"`
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)
	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

func (i *TodoItemUpdate) Validate() error {
	*i.Title = strings.TrimSpace(*i.Title)
	if *i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}
