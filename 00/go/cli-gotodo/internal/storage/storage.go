package gotodo_storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Pavlyysh/Develp10/00/go/cli-gotodo/internal/models"
)

func GetStoragePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".gotodo")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, "tasks.json"), nil
}

func LoadTasks() ([]models.Task, error) {
	path, err := GetStoragePath()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: failed to close file: %v\n", err)
		}
	}()

	var tasks []models.Task

	if info, _ := file.Stat(); info.Size() > 0 {
		if err := json.NewDecoder(file).Decode(&tasks); err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	path, err := GetStoragePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: failed to close file: %v\n", err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	return encoder.Encode(tasks)
}
