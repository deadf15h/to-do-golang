package menu

import (
	"bufio"
	"os"
	"strings"
	"to-do-golang/constants"
	"to-do-golang/utils"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		keyboardInput, _ := reader.ReadString('\n')
		keyboardInput = strings.TrimRight(keyboardInput, "\n")

		if keyboardInput[:2] == constants.KeyboardInputAdd {
			utils.CreateTask(utils.GetTaskValueFromInput(keyboardInput))
		}

		if keyboardInput == constants.KeyboardInputPrintAll {
			utils.PrintTasks()
		}

		if keyboardInput == constants.KeyboardInputSave {
			utils.SaveToJson()
		}

		if keyboardInput == constants.KeyboardInputLoad {
			utils.LoadFromJson()
		}

		if keyboardInput[:2] == constants.KeyboardInputChange {
			utils.ChangeTask(utils.GetTaskValueFromInput(keyboardInput))
		}

		if keyboardInput[:2] == constants.KeyboardInputDelete {
			utils.DeleteTask(utils.GetTaskValueFromInput(keyboardInput))
		}

		if keyboardInput == constants.KeyboardInputQuit {
			break
		}
	}
}
