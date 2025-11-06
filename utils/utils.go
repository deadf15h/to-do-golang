package utils

import (
	"errors"
	"fmt"
	"regexp"
	"to-do-golang/global"
)

func ParseTaskValueFromInput(input string) (string, error) {
	re := regexp.MustCompile(`"([^"]*)"`)
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) < 1 {
		return "", errors.New("error: incorrect input")
	}

	return matches[0][1], nil
}

func CreateTask(input string) {
	taskKey := input

	if taskKey == "" {
		return
	}

	if _, exists := global.Tasks[taskKey]; exists {
		fmt.Println("key already exists")
	} else {
		global.Tasks[taskKey] = false
	}

	PrintTasks()
}

func GetTaskValueFromInput(input string) string {
	matches, err := ParseTaskValueFromInput(input)

	if err == nil {
		return matches
	} else {
		return ""
	}
}

func PrintTasks() {
	for key, value := range global.Tasks {
		var printValue string

		if !value {
			printValue = "❌"
		} else {
			printValue = "✅"
		}

		fmt.Printf("%s %s\n", printValue, key)
	}
}

func ChangeTask(input string) {
	taskKey := input

	if taskKey == "" {
		return
	}

	if _, exists := global.Tasks[taskKey]; exists {
		global.Tasks[taskKey] = !global.Tasks[taskKey]

		PrintTasks()
	} else {
		fmt.Println("no key")
	}
}

func DeleteTask(input string) {
	taskKey := input

	if taskKey == "" {
		return
	}

	if _, exists := global.Tasks[taskKey]; exists {
		delete(global.Tasks, taskKey)
	} else {
		fmt.Println("no key")
	}

	PrintTasks()
}
