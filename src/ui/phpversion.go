package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetPHPVersion() string {
	reader := bufio.NewReader(os.Stdin)
	defaultVersion := "8.0"

	for {
		fmt.Printf("Enter a PHP version (must be > 8.0) [default %s]: ", defaultVersion)
		version, _ := reader.ReadString('\n')
		version = strings.TrimSpace(version)

		// if user pressed Enter, use default
		if version == "" {
			return defaultVersion
		}

		// Basic check: must start with "8.", "9.", or "10."
		if strings.HasPrefix(version, "8.") || strings.HasPrefix(version, "9.") || strings.HasPrefix(version, "10.") {
			return version
		}

		fmt.Println("Invalid version. Must be > 8.0, like 8.1 or 8.2")
	}
}
