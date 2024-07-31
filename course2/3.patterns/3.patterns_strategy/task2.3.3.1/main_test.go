package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define the mock services

type MockGistService struct {
	mock.Mock
}

func (m *MockGistService) List(ctx context.Context, username string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	args := m.Called(ctx, username, opts)
	var gists []*github.Gist
	if args.Get(0) != nil {
		gists = args.Get(0).([]*github.Gist)
	}
	return gists, args.Get(1).(*github.Response), args.Error(2)
}

type MockRepoService struct {
	mock.Mock
}

func (m *MockRepoService) List(ctx context.Context, username string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	args := m.Called(ctx, username, opts)
	var repos []*github.Repository
	if args.Get(0) != nil {
		repos = args.Get(0).([]*github.Repository)
	}
	return repos, args.Get(1).(*github.Response), args.Error(2)
}

func TestGithubGist_GetItems(t *testing.T) {
	mockService := new(MockGistService)
	ctx := context.Background()
	username := "testuser"

	gistID := "gist1"
	description := "Gist description"
	htmlURL := "https://gist.github.com/gist1"

	gists := []*github.Gist{
		{
			ID:          &gistID,
			Description: &description,
			HTMLURL:     &htmlURL,
		},
	}

	mockService.On("List", ctx, username, (*github.GistListOptions)(nil)).Return(gists, &github.Response{}, nil)

	gist := NewGithubGist(mockService)
	items, err := gist.GetItems(ctx, username)

	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "gist1", items[0].Title)
	assert.Equal(t, "Gist description", items[0].Описание)
	assert.Equal(t, "https://gist.github.com/gist1", items[0].Link)

	mockService.AssertExpectations(t)
}

func TestGithubGist_GetItems_Error(t *testing.T) {
	mockService := new(MockGistService)
	ctx := context.Background()
	username := "testuser"

	mockService.On("List", ctx, username, (*github.GistListOptions)(nil)).Return(nil, &github.Response{}, assert.AnError)

	gist := NewGithubGist(mockService)
	items, err := gist.GetItems(ctx, username)

	assert.Error(t, err)
	assert.Nil(t, items)

	mockService.AssertExpectations(t)
}

func TestGithubRepo_GetItems(t *testing.T) {
	mockService := new(MockRepoService)
	ctx := context.Background()
	username := "testuser"

	repoName := "repo1"
	description := "Repo description"
	htmlURL := "https://github.com/repo1"

	repos := []*github.Repository{
		{
			Name:        &repoName,
			Description: &description,
			HTMLURL:     &htmlURL,
		},
	}

	mockService.On("List", ctx, username, (*github.RepositoryListOptions)(nil)).Return(repos, &github.Response{}, nil)

	repo := NewGithubRepo(mockService)
	items, err := repo.GetItems(ctx, username)

	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "repo1", items[0].Title)
	assert.Equal(t, "Repo description", items[0].Описание)
	assert.Equal(t, "https://github.com/repo1", items[0].Link)

	mockService.AssertExpectations(t)
}

func TestGithubRepo_GetItems_Error(t *testing.T) {
	mockService := new(MockRepoService)
	ctx := context.Background()
	username := "testuser"

	mockService.On("List", ctx, username, (*github.RepositoryListOptions)(nil)).Return(nil, &github.Response{}, assert.AnError)

	repo := NewGithubRepo(mockService)
	items, err := repo.GetItems(ctx, username)

	assert.Error(t, err)
	assert.Nil(t, items)

	mockService.AssertExpectations(t)
}

func TestRealGeneralGithubLister_GetItems(t *testing.T) {
	mockGistService := new(MockGistService)
	mockRepoService := new(MockRepoService)
	ctx := context.Background()
	username := "testuser"

	gistID := "gist1"
	gistDescription := "Gist description"
	gistHTMLURL := "https://gist.github.com/gist1"

	gists := []*github.Gist{
		{
			ID:          &gistID,
			Description: &gistDescription,
			HTMLURL:     &gistHTMLURL,
		},
	}

	repoName := "repo1"
	repoDescription := "Repo description"
	repoHTMLURL := "https://github.com/repo1"

	repos := []*github.Repository{
		{
			Name:        &repoName,
			Description: &repoDescription,
			HTMLURL:     &repoHTMLURL,
		},
	}

	mockGistService.On("List", ctx, username, (*github.GistListOptions)(nil)).Return(gists, &github.Response{}, nil)
	mockRepoService.On("List", ctx, username, (*github.RepositoryListOptions)(nil)).Return(repos, &github.Response{}, nil)

	gist := NewGithubGist(mockGistService)
	repo := NewGithubRepo(mockRepoService)
	gg := NewGeneralGithub()

	items, err := gg.GetItems(ctx, username, gist)
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "gist1", items[0].Title)
	assert.Equal(t, "Gist description", items[0].Описание)
	assert.Equal(t, "https://gist.github.com/gist1", items[0].Link)

	items, err = gg.GetItems(ctx, username, repo)
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "repo1", items[0].Title)
	assert.Equal(t, "Repo description", items[0].Описание)
	assert.Equal(t, "https://github.com/repo1", items[0].Link)

	mockGistService.AssertExpectations(t)
	mockRepoService.AssertExpectations(t)
}
