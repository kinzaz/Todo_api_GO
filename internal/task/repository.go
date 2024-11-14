package task

import "todo_GO/pkg/db"

type TaskRepository struct {
	Database *db.DB
}

func NewTaskRepository(database *db.DB) *TaskRepository {
	return &TaskRepository{
		Database: database,
	}
}

func (repo *TaskRepository) Create(task *Task) (*Task, error) {
	result := repo.Database.DB.Create(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}
