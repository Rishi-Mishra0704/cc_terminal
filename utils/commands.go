package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
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

func Echo(args []string) string {
	return strings.Join(args, " ")
}

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Error getting current directory: %s", err)
	}
	return dir
}

func Cat(args []string) string {
	if len(args) != 1 {
		return "Usage: cat <file>"
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}
	return string(data)
}

func Touch(args []string) string {
	if len(args) != 1 {
		return "Usage: touch <file>"
	}
	file, err := os.Create(args[0])
	if err != nil {
		return fmt.Sprintf("Error creating file: %s", err)
	}
	file.Close()
	return "File created successfully"
}

func Clear() string {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return ""
}

func Date() string {
	return time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
}
