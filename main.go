package main

import (
	"log"
	"phpvm/src/services/bootstrap"
	"phpvm/src/ui"
)

func main() {
	projectName := ui.GetProjectName()
	phpVersion := ui.GetPHPVersion()
	projectPath, _ := bootstrap.ProjectPath(projectName)

	// کپی template ها
	bootstrap.CopyTheTemplate(projectPath)

	// ویرایش فایل .env
	bootstrap.UpdateEnvFile(projectPath, projectName, phpVersion)
	if bootstrap.CheckDockerInstalled() {
		bootstrap.ShowShellFunctionsMessage()
		log.Println("Docker and Docker Compose detected. Starting containers...")
		bootstrap.RunDockerCompose(projectPath)
		
	} else {
		log.Println("Docker or docker-compose not found. Please install them to run containers.")
	}

}
