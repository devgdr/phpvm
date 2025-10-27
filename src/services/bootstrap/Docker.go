package bootstrap

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

// CheckDockerInstalled بررسی می‌کند Docker و Docker Compose نصب هستند
func CheckDockerInstalled() bool {
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Println("Docker is not installed")
		return false
	}

	_, err = exec.LookPath("docker-compose")
	if err != nil {
		log.Println("docker-compose is not installed")
		return false
	}

	return true
}

// RunDockerCompose runs "docker-compose up -d" inside the project folder
// and shows real-time output and errors.
func RunDockerCompose(projectPath string) {
	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Dir = projectPath  // مسیر پروژه
	cmd.Stdout = os.Stdout // نمایش خروجی
	cmd.Stderr = os.Stderr // نمایش خطاها

	log.Println("🟢 Running 'docker-compose up -d' in", projectPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("❌ Error running docker-compose: %v", err)
	}

	log.Println("✅ Docker Compose started successfully")
}

func ShowShellFunctionsMessage() {
	if runtime.GOOS == "windows" {
		log.Println("Add the following functions to your PowerShell profile ($PROFILE):Remmber the only acssessable under the project directory")
		log.Println(`
function php { docker-compose exec -w /var/www/html -u developer php php $args }
function composer { docker-compose exec -w /var/www/html -u developer php composer $args }
function composerexec { docker-compose exec -w /var/www/html -u developer php composer exec $args }
`)
	} else {
		log.Println("Add the following functions to your shell profile (~/.bashrc or ~/.zshrc or whatever u are using!):")
		log.Println(`
php() {
    docker-compose exec -w /var/www/html -u developer php php "$@"
}
composer() {
    docker-compose exec -w /var/www/html -u developer php composer "$@"
}
composerexec() {
    docker-compose exec -w /var/www/html -u developer php composer exec "$@"
}
Then run 'source ~/.bashrc' (or 'source ~/.zshrc') to use them.
`)
	}
}
