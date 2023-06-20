package service

import "github.com/phucnh/go-app-sample/entity"

// TaskService defines method for task service interface
type TaskService interface {
	CreateTask(*entity.Task) (*entity.Task, error)
	GetTasksList(limit int, offset int) ([]*entity.Task, int64, error)
}
