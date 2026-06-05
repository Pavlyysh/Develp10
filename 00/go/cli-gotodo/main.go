package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

var (
	ErrEmptyTitle   = errors.New("title cannot be empty")
	ErrTaskNotFound = errors.New("task not found")
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Due         string `json:"due"`
}

// CLI-утилита gotodo — менеджер задач в терминале.

// Критерии приёмки:

// Команды: add, list, done, rm, clear.
// Хранение в JSON-файле в ~/.gotodo/tasks.json.
// Цветной вывод (через fatih/color или ANSI-коды).
// Флаги: --priority, --due, --filter.
// Свой репозиторий на GitHub с README, MIT-лицензией, минимум 5 осмысленных коммитов.
// Установка через go install.
// golangci-lint чист.

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotodo <option>")
		fmt.Println("Example: gotodo add -t PetProject -d Make gotodo petProject")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("t", "", "title")
		desc := addCmd.String("d", "", "description")
		priority := addCmd.String("priority", "medium", "priority, medium by default")
		due := addCmd.String("due", "", "due date")
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
		task := &Task{
			Title:       *title,
			Description: *desc,
			Priority:    *priority,
			Due:         *due,
			Status:      "pending",
		}
		err = add(task)
		if err != nil {
			fmt.Println(err)
		}
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "all", "filter tasks: all, pending, done, overdue")
		listCmd.Parse(os.Args[2:])

		err := list(*filter)
		if err != nil {
			fmt.Println(err)
		}
	case "rm":
		rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
		title := rmCmd.String("t", "", "title")
		err := rmCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}

		task := &Task{Title: *title}
		err = rm(task)
		if err != nil {
			fmt.Println(err)
		}
	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		title := doneCmd.String("t", "", "title")
		err := doneCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
		}
		task := &Task{Title: *title}

		err = done(task)
		if err != nil {
			fmt.Println(err)
		}
	case "clear":
		if err := clear(); err != nil {
			fmt.Println(err)
		}
	default:
	}

}

func add(task *Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task.ID = id

	tasks = append(tasks, *task)

	return saveTasks(tasks)
}

func list(filter string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	filtered := filterTasks(tasks, filter)

	if len(filtered) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for _, task := range filtered {
		printTask(task)
	}

	return nil
}
func filterTasks(tasks []Task, filter string) []Task {
	switch filter {
	case "pending":
		return filterByStatus(tasks, "pending")
	case "done":
		return filterByStatus(tasks, "done")
	case "overdue":
		return filterOverdue(tasks)
	case "all", "":
		return tasks
	default:
		fmt.Printf("Unknown filter: %s, showing all tasks\n", filter)
		return tasks
	}
}

func filterByStatus(tasks []Task, status string) []Task {
	var result []Task
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}

func filterOverdue(tasks []Task) []Task {
	var result []Task
	today := time.Now().Format("2006-01-02")

	for _, task := range tasks {
		if task.Status != "done" && task.Due != "" && task.Due < today {
			result = append(result, task)
		}
	}
	return result
}

func printTask(task Task) {
	switch task.Status {
	case "done":
		color.Green("[✓] %s", task.Title)
	case "pending":
		color.Yellow("[ ] %s", task.Title)
	}

	if task.Description != "" {
		fmt.Printf("    %s\n", task.Description)
	}
	if task.Due != "" {
		fmt.Printf("    Due: %s\n", task.Due)
	}
	if task.Priority != "" {
		fmt.Printf("    Priority: %s\n", task.Priority)
	}
	fmt.Println("------------------")
}

func done(inputTask *Task) error {
	if inputTask.Title == "" {
		return ErrEmptyTitle
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Title == inputTask.Title {
			tasks[i].Status = "done"
			if err := saveTasks(tasks); err != nil {
				return err
			}
			return nil
		}

	}

	return ErrTaskNotFound
}

func rm(inputTask *Task) error {
	if inputTask.Title == "" {
		return ErrEmptyTitle
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Title == inputTask.Title {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := saveTasks(tasks); err != nil {
				return err
			}
			return nil
		}
	}
	return ErrTaskNotFound
}

func clear() error {
	tasks := []Task{}

	return saveTasks(tasks)
}

func getStoragePath() (string, error) {
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

func loadTasks() ([]Task, error) {
	path, err := getStoragePath()
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

	var tasks []Task

	if info, _ := file.Stat(); info.Size() > 0 {
		if err := json.NewDecoder(file).Decode(&tasks); err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	path, err := getStoragePath()
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
