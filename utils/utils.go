package utils

import (
	"fmt"
	"regexp"
	"to-do-golang/global"
)

func ParseTaskValueFromInput(input string) [][]string {
	re := regexp.MustCompile(`"([^"]*)"`)
	matches := re.FindAllStringSubmatch(input, -1)
	return matches
}

func GetTaskValueFromInput(input string) string {
	matches := ParseTaskValueFromInput(input)

	for _, match := range matches {
		if len(match) > 1 {
			global.Tasks[match[0]] = false
		}
	}

	return input
}

func PrintTasks() {
	for key, value := range global.Tasks {
		var printValue string

		if value {
			printValue = "[X]"
		} else {
			printValue = "[O]"
		}

		fmt.Printf("%s %s\n", printValue, key)
	}
}

func ChangeTask(input string) {
	taskKey := ParseTaskValueFromInput(input)[0][0]
	_, found := global.Tasks[taskKey]

	if !found {
		// TODO add err
		fmt.Println("no key")
	} else {
		// TODO only to true???
		global.Tasks[taskKey] = !global.Tasks[taskKey]

		PrintTasks()
	}
}
