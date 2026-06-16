package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Pavlyysh/Develp10/module2/go/gotodo-taskStore-interface/internal/models"
)

type JSONStore struct {
	path string
}

func NewJSONStore() (*JSONStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get home dir: %w", err)
	}

	dir := filepath.Join(home, ".gotodo")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed create dir: %w", err)
	}

	return &JSONStore{
		path: filepath.Join(dir, "tasks.json"),
	}, nil
}

func (s *JSONStore) Load() ([]models.Task, error) {
	file, err := os.OpenFile(s.path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("file open error: %w", err)
	}
	defer file.Close()

	var tasks []models.Task
	json.NewDecoder()
}

func (s *JSONStore) Save(tasks []models.Task) error {}
