package git

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	giturl "github.com/kubescape/go-git-url"
)

const DefaultRemote = "origin"

var (
	// errors
	ErrNoRemoteURLs = errors.New("no remote URLs found")
)

func GetCurrentRepoSlug() (string, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return "", fmt.Errorf("could not find git repo, repo may not be initialised: %w", err)
	}

	rem, err := repo.Remote(DefaultRemote)
	if err != nil {
		return "", fmt.Errorf("could not get remote %s for repo: %w", DefaultRemote, err)
	}

	gitURLs := rem.Config().URLs

	if len(gitURLs) == 0 {
		return "", ErrNoRemoteURLs
	}

	return getSlugFromURL(gitURLs[0])
}

func getSlugFromURL(gitURL string) (string, error) {
	parsedURL, err := giturl.NewGitURL(gitURL)
	if err != nil {
		return "", fmt.Errorf("error parsing git URL %q: %w", gitURL, err)
	}

	return strings.Join([]string{parsedURL.GetOwnerName(), parsedURL.GetRepoName()}, "/"), nil
}
