package repository

import (
	"database/sql"
	"time"

	"github.com/EvoSched/gotask/internal/models"
)

var due = time.Now()
var due1 = time.Now().AddDate(0, 0, 2)
var due2 = time.Now().AddDate(0, 0, 4)

// sample data to test command functions
var tasks = []*models.Task{
	models.NewTask(1, "description1", &due, []string{"MA"}),
	models.NewTask(2, "description2", &due1, []string{"CS"}),
	models.NewTask(3, "description3", &due2, []string{"MA"}),
	models.NewTask(4, "description4", nil, []string{"CS"}),
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) GetTask(id int) (*models.Task, error) {
	//TODO: sql query

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, nil
}

func (r *TaskRepository) GetTasks() ([]*models.Task, error) {
	//TODO: sql query

	return tasks, nil
}
