package main

import (
	"github.com/jonascavalcanti/devpie/internal/app/github"
	"github.com/jonascavalcanti/devpie/internal/delivery/http"
)

func main() {
	githubRepo := github.NewGitHubRepository()
	handler := http.NewHandler(githubRepo)

	handler.Run(":8080")
}
