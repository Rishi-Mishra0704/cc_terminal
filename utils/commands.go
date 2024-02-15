package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ListFiles() string {
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Sprintf("Error listing files: %s", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return strings.Join(fileNames, "\n")
}

func CreateDirectory(args []string) string {
	if len(args) != 1 {
		return "Usage: mkdir <directory>"
	}
	err := os.Mkdir(args[0], 0755)
	if err != nil {
		return fmt.Sprintf("Error creating directory: %s", err)
	}
	return "Directory created successfully"
}

func ChangeDirectory(args []string) string {
	if len(args) != 1 {
		return "Usage: cd <directory>"
	}
	err := os.Chdir(args[0])
	if err != nil {
		return fmt.Sprintf("Error changing directory: %s", err)
	}
	return "Directory changed successfully"
}

func ExecuteCommand(command string, args []string) string {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing command: %s", err)
	}
	return string(out)
}
