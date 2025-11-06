package utils

import (
	"fmt"
	"regexp"
	"to-do-golang/global"
)

func ParseTaskValueFromInput(input string) string {
	re := regexp.MustCompile(`"([^"]*)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	return matches[0][1]
}

func CreateTask(taskKey string) {
	// TODO add err
	if _, exists := global.Tasks[taskKey]; exists {
		fmt.Println("key already exists")
	} else {
		global.Tasks[taskKey] = false
	}

	PrintTasks()
}

func GetTaskValueFromInput(input string) string {
	matches := ParseTaskValueFromInput(input)

	return matches
}

func PrintTasks() {
	for key, value := range global.Tasks {
		var printValue string

		if !value {
			printValue = "[O]"
		} else {
			printValue = "[X]"
		}

		fmt.Printf("%s %s\n", printValue, key)
	}
}

func ChangeTask(input string) {
	taskKey := input

	if _, exists := global.Tasks[taskKey]; exists {
		global.Tasks[taskKey] = !global.Tasks[taskKey]

		PrintTasks()
	} else {
		// TODO add err
		fmt.Println("no key")
	}
}

func DeleteTask(input string) {
	taskKey := input

	if _, exists := global.Tasks[taskKey]; exists {
		// TODO add err
		delete(global.Tasks, taskKey)
	} else {
		// TODO add err
		fmt.Println("no key")
	}

	PrintTasks()
}
