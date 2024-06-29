package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

// RepoLister интерфейс для получения списка репозиториев пользователя
type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

// GistLister интерфейс для получения списка гистов пользователя
type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

// Githuber интерфейс, который должен быть реализован адаптером
type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

// GithubAdapter адаптер для обеспечения совместимости между интерфейсом Githuber и библиотекой github.com/google/go-github/v53/github
type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

// NewGithubAdapter функция для создания нового экземпляра GithubAdapter
func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	return &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}
}

// Item структура для представления элемента (репозиторий или гист)
type Item struct {
	Title       string
	Description string
	Link        string
}

// GetGists метод адаптера для получения списка гистов пользователя
func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, gist := range gists {
		item := Item{
			Title:       gist.GetDescription(),
			Description: gist.GetDescription(),
			Link:        gist.GetHTMLURL(),
		}
		items = append(items, item)
	}
	return items, nil
}

// GetRepos метод адаптера для получения списка репозиториев пользователя
func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, repo := range repos {
		item := Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetHTMLURL(),
		}
		items = append(items, item)
	}
	return items, nil
}

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewGithubProxy(github Githuber) *GithubProxy {
	return &GithubProxy{
		github: github,
		cache:  make(map[string][]Item),
	}
}

// GetGists метод прокси для получения/кеширования списка гистов пользователя
func (g *GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	if _, ok := g.cache[fmt.Sprintf(username+"_gists")]; !ok {
		gists, err := g.github.GetGists(ctx, username)
		if err != nil {
			return nil, err
		}

		var items []Item
		for _, gist := range gists {
			item := Item{
				Title:       gist.Title,
				Description: gist.Description,
				Link:        gist.Link,
			}
			items = append(items, item)
		}
		g.cache[fmt.Sprintf(username+"_gists")] = items
		println("got gists from web")
		return items, nil
	}
	println("got gists from cache")
	return g.cache[fmt.Sprintf(username+"_gists")], nil
}

// GetRepos метод прокси для получения/кеширования списка репозиториев пользователя
func (g *GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	if _, ok := g.cache[fmt.Sprintf(username+"_repos")]; !ok {
		repos, err := g.github.GetRepos(ctx, username)
		if err != nil {
			return nil, err
		}

		var items []Item
		for _, repo := range repos {
			item := Item{
				Title:       repo.Title,
				Description: repo.Description,
				Link:        repo.Link,
			}
			items = append(items, item)
		}
		g.cache[fmt.Sprintf(username+"_repos")] = items
		println("got repos from web")
		return items, nil
	}
	println("got repos from cache")
	return g.cache[fmt.Sprintf(username+"_repos")], nil
}

func main() {
	ctx := context.Background()

	// Получаем токен из переменной окружения
	token := os.Getenv("GITHUB_PAT")
	if token == "" {
		log.Fatal("GITHUB_PAT environment variable is not set")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	g := NewGithubAdapter(client)

	proxy := NewGithubProxy(g)

	// Тест кэширования
	for i := 0; i < 3; i++ {
		// Пример использования прокси для получения гистов
		gists, err := proxy.GetGists(ctx, "ptflp")
		if err != nil {
			log.Fatalf("Error fetching gists: %v", err)
		}
		fmt.Println("Gists:")
		for _, gist := range gists {
			fmt.Printf("- %s: %s (%s)\n", gist.Title, gist.Description, gist.Link)
		}

		// Пример использования прокси для получения репозиториев
		repos, err := proxy.GetRepos(ctx, "ptflp")
		if err != nil {
			log.Fatalf("Error fetching repos: %v", err)
		}
		fmt.Println("Repositories:")
		for _, repo := range repos {
			fmt.Printf("- %s: %s (%s)\n", repo.Title, repo.Description, repo.Link)
		}
	}
}
