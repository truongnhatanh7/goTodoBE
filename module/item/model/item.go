package model

import (
	"errors"
	"strings"

	"github.com/truongnhatanh7/goTodoBE/common"
)

var (
  ErrTitleCannotBeEmpty = errors.New("Title cannot be empty")
)

type TodoItem struct {
  common.SQLModel
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
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
