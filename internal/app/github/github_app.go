package github

type GitHubRepository interface {
	GetRepositories() ([]string, error)
}

type app struct {
	githubRepo GitHubRepository
}

func (a *app) GetRepositories() ([]string, error) {
	return a.githubRepo.GetRepositories()
}
