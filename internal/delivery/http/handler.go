package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonascavalcanti/devpie/internal/app/github"
)

type Handler struct {
	gitHubRepo github.GitHubRepository
}

func NewHandler(githubRepository github.GitHubRepository) *Handler {
	return &Handler{gitHubRepo: githubRepository}
}

func (h *Handler) GetRepositories(c *gin.Context) {
	repos, err := h.gitHubRepo.GetRepositories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"repositories": repos})
}

func (h *Handler) Run(address string) {
	router := gin.Default()
	router.GET("/repositories", h.GetRepositories)

	router.Run(address)
}
