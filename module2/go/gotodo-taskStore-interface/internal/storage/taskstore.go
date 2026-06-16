package storage

import "github.com/Pavlyysh/Develp10/module2/go/gotodo-taskStore-interface/internal/models"

type TaskStore interface {
	Load() ([]models.Task, error)
	Store(tasks []models.Task) error
}
