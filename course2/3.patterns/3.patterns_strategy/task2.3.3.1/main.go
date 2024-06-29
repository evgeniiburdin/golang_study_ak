package main

import (
	"context"
	"github.com/google/go-github/v53/github"
)

type Item struct {
	Title    string
	Описание string
	Link     string
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GistService interface {
	List(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type RepoService interface {
	List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GithubGist struct {
	service GistService
}

func NewGithubGist(service GistService) *GithubGist {
	return &GithubGist{service: service}
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.service.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, item := range gists {
		title := "<no title>"
		desc := "<no description>"
		link := "<no link>"
		if item.ID != nil {
			title = *item.ID
		}
		if item.Description != nil {
			desc = *item.Description
		}
		if item.HTMLURL != nil {
			link = *item.HTMLURL
		}
		items = append(items, Item{
			Title:    title,
			Описание: desc,
			Link:     link,
		})
	}
	return items, nil
}

type GithubRepo struct {
	service RepoService
}

func NewGithubRepo(service RepoService) *GithubRepo {
	return &GithubRepo{service: service}
}

func (g *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.service.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, item := range repos {
		title := "<no title>"
		desc := "<no description>"
		link := "<no link>"
		if item.Name != nil {
			title = *item.Name
		}
		if item.Description != nil {
			desc = *item.Description
		}
		if item.HTMLURL != nil {
			link = *item.HTMLURL
		}
		items = append(items, Item{
			Title:    title,
			Описание: desc,
			Link:     link,
		})
	}
	return items, nil
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type RealGeneralGithubLister struct{}

func (g *RealGeneralGithubLister) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

func NewGeneralGithub() GeneralGithubLister {
	return &RealGeneralGithubLister{}
}

/*
func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_PAT")})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client.Gists)
	repo := NewGithubRepo(client.Repositories)

	gg := NewGeneralGithub()

	data, err := gg.GetItems(ctx, "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("GISTS:")
	for _, item := range data {
		fmt.Println(item)
	}

	fmt.Println()

	data, err = gg.GetItems(ctx, "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("REPOS:")
	for _, item := range data {
		fmt.Println(item)
	}
}
*/
