package gotodo_task

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Pavlyysh/Develp10/00/go/cli-gotodo/internal/models"
	gstorage "github.com/Pavlyysh/Develp10/00/go/cli-gotodo/internal/storage"
	"github.com/fatih/color"
)

var (
	ErrEmptyTitle   = errors.New("title cannot be empty")
	ErrTaskNotFound = errors.New("task not found")
)

func add(task *models.Task) error {
	if task.Title == "" {
		return ErrEmptyTitle
	}
	tasks, err := gstorage.LoadTasks()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task.ID = id

	tasks = append(tasks, *task)

	return gstorage.SaveTasks(tasks)
}

func list(filter string) error {
	tasks, err := gstorage.LoadTasks()
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
func filterTasks(tasks []models.Task, filter string) []models.Task {
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

func filterByStatus(tasks []models.Task, status string) []models.Task {
	var result []models.Task
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}

func filterOverdue(tasks []models.Task) []models.Task {
	var result []models.Task
	today := time.Now().Format("2006-01-02")

	for _, task := range tasks {
		if task.Status != "done" && task.Due != "" && task.Due < today {
			result = append(result, task)
		}
	}
	return result
}

func printTask(task models.Task) {
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

func done(inputTask *models.Task) error {
	if inputTask.Title == "" {
		return ErrEmptyTitle
	}

	tasks, err := gstorage.LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Title == inputTask.Title {
			tasks[i].Status = "done"
			if err := gstorage.SaveTasks(tasks); err != nil {
				return err
			}
			return nil
		}

	}

	return ErrTaskNotFound
}

func rm(inputTask *models.Task) error {
	if inputTask.Title == "" {
		return ErrEmptyTitle
	}

	tasks, err := gstorage.LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Title == inputTask.Title {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := gstorage.SaveTasks(tasks); err != nil {
				return err
			}
			return nil
		}
	}
	return ErrTaskNotFound
}

func clear() error {
	tasks := []models.Task{}

	return gstorage.SaveTasks(tasks)
}

func Match(s string) error {
	switch s {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("t", "", "title")
		desc := addCmd.String("d", "", "description")
		priority := addCmd.String("priority", "medium", "priority, medium by default")
		due := addCmd.String("due", "", "due date")
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		task := &models.Task{
			Title:       *title,
			Description: *desc,
			Priority:    *priority,
			Due:         *due,
			Status:      "pending",
		}
		err = add(task)
		if err != nil {
			return err
		}
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "all", "filter tasks: all, pending, done, overdue")
		listCmd.Parse(os.Args[2:])

		err := list(*filter)
		if err != nil {
			return err
		}
	case "rm":
		rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
		title := rmCmd.String("t", "", "title")
		err := rmCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}

		task := &models.Task{Title: *title}
		err = rm(task)
		if err != nil {
			return err
		}
	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		title := doneCmd.String("t", "", "title")
		err := doneCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		task := &models.Task{Title: *title}

		err = done(task)
		if err != nil {
			return err
		}
	case "clear":
		if err := clear(); err != nil {
			return err
		}
	default:
		return errors.New("method not allowed")
	}

	return nil
}
