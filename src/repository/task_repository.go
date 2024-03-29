package repository

import "github.com/phucnh/go-app-sample/entity"

type TaskRepository interface {
	CreateTask(task *entity.Task) (*entity.Task, error)
	GetTasksList(limit int, offset int) ([]*entity.Task, int64, error)
}
