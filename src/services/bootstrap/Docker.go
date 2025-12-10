package bootstrap

import (
	"log"
	"os"
	"os/exec"
)

// CheckDockerInstalled بررسی می‌کند Docker و Docker Compose نصب هستند
func CheckDockerInstalled() bool {
	// Check if docker is installed
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Println("Docker is not installed")
		return false
	}

	// Check if docker compose subcommand works
	cmd := exec.Command("docker", "compose", "version")
	err = cmd.Run()
	if err != nil {
		log.Println("docker compose is not installed")
		return false
	}

	return true
}

// RunDockerCompose runs "docker-compose up -d" inside the project folder
// and shows real-time output and errors.
func RunDockerCompose(projectPath string) {
	// Correct: separate "docker" and "compose" as individual arguments
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = projectPath  // مسیر پروژه
	cmd.Stdout = os.Stdout // نمایش خروجی
	cmd.Stderr = os.Stderr // نمایش خطاها

	log.Println("🟢 Running 'docker compose up -d' in", projectPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("❌ Error running docker-compose: %v", err)
	}

	log.Println("✅ Docker Compose started successfully")
}
func ShowUsageInstructions() {
	log.Println("\nSetup complete! You can now use the following 'make' commands:")
	log.Println("  make up              # Start containers")
	log.Println("  make down            # Stop containers")
	log.Println("  make php [cmd]       # Run PHP commands (e.g., make php -v)")
	log.Println("  make composer [cmd]  # Run Composer commands (e.g., make composer install)")
	log.Println("  make logs            # View logs")
	log.Println("\nMake sure you have 'make' installed on your system.")
}
