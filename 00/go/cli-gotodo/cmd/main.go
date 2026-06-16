package main

import (
	"fmt"
	"os"

	gtask "github.com/Pavlyysh/Develp10/00/go/cli-gotodo/internal/task"
)

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

	if err := gtask.Match(os.Args[1]); err != nil {
		fmt.Println(err)
	}

}
