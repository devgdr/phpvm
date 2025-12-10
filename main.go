package main

import (
	"log"
	"phpvm/src/services/bootstrap"
	"phpvm/src/ui"
)

func main() {
	projectName := ui.GetProjectName()
	phpVersion := ui.GetPHPVersion()
	projectPath, err := bootstrap.ProjectPath(projectName)
	if err != nil {
		log.Fatalf("Error resolving project path: %v", err)
	}


	bootstrap.CopyTheTemplate(projectPath)


	bootstrap.UpdateEnvFile(projectPath, projectName, phpVersion)
	if bootstrap.CheckDockerInstalled() {
		bootstrap.ShowUsageInstructions()
		log.Println("Docker and Docker Compose detected. Starting containers...")
		bootstrap.RunDockerCompose(projectPath)

	} else {
		log.Println("Docker or docker-compose not found. Please install them to run containers.")
	}

}
