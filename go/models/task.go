package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Task "Object
type Task struct {
	//ID        uuid.UUID `json:"id"`
	ID        string    `json:"id"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool      `json:"completed"`
}

// 定义 Priority 类型，对应 PostgreSQL 中的 "Priority" 类型
type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

// func (Todo) TableName() string {
// 	return "Todo"
// }

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
