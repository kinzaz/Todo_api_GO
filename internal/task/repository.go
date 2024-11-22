package task

import (
	"todo_GO/pkg/db"
)

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

func (repo *TaskRepository) Delete(id string) error {
	result := repo.Database.DB.Delete(&Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TaskRepository) GetById(id uint) (*Task, error) {
	var task Task
	result := repo.Database.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (repo *TaskRepository) Update(task *Task) (*Task, error) {
	result := repo.Database.DB.Save(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	result := repo.Database.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
