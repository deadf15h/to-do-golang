package main

import (
	"testing"
	"to-do-golang/utils"
)

func TestParseTaskValueFromInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal input",
			input:    `a: "simple task"`,
			expected: "simple task",
		},
		{
			name:     "empty input",
			input:    `:a ""`,
			expected: "",
		},
		{
			name:     "multiple input, only first value taken",
			input:    `:a "task1" "task2"`,
			expected: "task1",
		},
		{
			name:     "single quot mark left input error",
			input:    `:a "task1`,
			expected: "",
		},
		{
			name:     "single quot mark right input error",
			input:    `:a task1"`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := utils.ParseTaskValueFromInput(tt.input)

			if result != tt.expected {
				t.Errorf("ParseTaskValueFromInput(%q) = %q, expected %q\n",
					tt.input, result, tt.expected)
			}
		})
	}
}
