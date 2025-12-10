package bootstrap

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// UpdateEnvFile تغییر مقادیر PHP_BASE و PROJECT_NAME در فایل .env
func UpdateEnvFile(projectPath, projectName, phpVersion string) {
	envPath := filepath.Join(projectPath, ".env")

	// فایل .env را بخوان
	data, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	content := string(data)

	// مقادیر مورد نظر را جایگزین کن
	content = replaceEnvValue(content, "PHP_BASE", "php:"+phpVersion+"-fpm")
	content = replaceEnvValue(content, "PROJECT_NAME", projectName)

	// فایل .env را دوباره بنویس
	if err := os.WriteFile(envPath, []byte(content), 0644); err != nil {
		log.Fatalf("Error writing .env file: %v", err)
	}

	log.Println("✅ .env updated successfully")
}

// تابع کمکی برای جایگزینی مقدار یک کلید در فایل env
func replaceEnvValue(content, key, value string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = key + "=" + value
		}
	}
	return strings.Join(lines, "\n")
}
