package scheduler

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Duration string
	Error    string
	Name     string
	gorm.Model
}

type Scheduler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Scheduler {
	return Scheduler{db}
}

func (s Scheduler) Do(c context.Context, t *Task, f func() error) {
	s.DB.WithContext(c).Create(t)
	defer func() {
		t.Duration = time.Since(t.CreatedAt).Truncate(time.Second).String()
		s.DB.WithContext(c).Updates(t)
	}()
	err := f()
	if err != nil {
		t.Error = err.Error()
	}
}
