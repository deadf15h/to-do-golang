package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"to-do-golang/constants"
	"to-do-golang/global"
)

func ParseTaskValueFromInput(input string) (string, error) {
	re := regexp.MustCompile(`"([^"]*)"`)
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) < 1 {
		return "", errors.New(constants.ErrorIncorrectInputMsg)
	}

	return matches[0][1], nil
}

func CreateTask(input string) {
	taskKey := input

	if taskKey == "" {
		return
	}

	if _, exists := global.Tasks[taskKey]; exists {
		fmt.Println(constants.KeyAlreadyExistsMsg)
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
			printValue = constants.PrintValueFalse
		} else {
			printValue = constants.PrintValueTrue
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
		fmt.Println(constants.NoKeyMsg)
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
		fmt.Println(constants.NoKeyMsg)
	}

	PrintTasks()
}

func SaveToJson() {
	jsonData, err := json.MarshalIndent(global.Tasks, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("data.json", jsonData, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func LoadFromJson() {
	data, err := os.ReadFile("data.json")

	if err != nil {
		log.Fatal(constants.ErrorReadingFileMsg, err)
	}

	var res map[string]bool

	err = json.Unmarshal(data, &res)

	if err != nil {
		log.Fatal(constants.ErrorParseFileMsg, err)
	}

	if err == nil {
		fmt.Println(constants.LoadingFromFileSuccessMsg)
	}

	global.Tasks = res

	PrintTasks()
}
