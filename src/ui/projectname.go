package ui

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetProjectName() string {
	reader := bufio.NewReader(os.Stdin)

	validName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

	for {
		fmt.Print("Enter project name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Project name cannot be empty")
			continue
		}

		if !validName.MatchString(input) {
			fmt.Println("Invalid project name. Must start with a letter and can contain letters, numbers, '-' or '_' only")
			continue
		}

		return strings.ToLower(input)
	}
}

