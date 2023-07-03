package main

import (
	"encoding/json"
	"log"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	now := time.Now().UTC()
	item := TodoItem{
		Id:          1,
		Title:       "Task 1",
		Description: "Content 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

  jsData, err := json.Marshal(item)
  
  if err != nil {
    log.Fatalln(err)
  }

  log.Println(string(jsData))
}
