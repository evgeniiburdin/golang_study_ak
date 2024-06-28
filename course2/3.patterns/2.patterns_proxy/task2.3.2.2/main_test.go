package main

import (
	"context"
	"github.com/google/go-github/v53/github"
	"reflect"
	"testing"
)

// MockRepoLister является мок-реализацией интерфейса RepoLister
type MockRepoLister struct{}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	repos := []*github.Repository{
		{
			Name:        github.String("Repo1"),
			Description: github.String("Description1"),
			HTMLURL:     github.String("http://github.com/user/repo1"),
		},
		{
			Name:        github.String("Repo2"),
			Description: github.String("Description2"),
			HTMLURL:     github.String("http://github.com/user/repo2"),
		},
		{
			Name:        github.String("Repo3"),
			Description: github.String("Description3"),
			HTMLURL:     github.String("http://github.com/user/repo3"),
		},
		{
			Name:        github.String("Repo4"),
			Description: github.String("Description4"),
			HTMLURL:     github.String("http://github.com/user/repo4"),
		},
	}
	return repos, nil, nil
}

// MockGistLister является мок-реализацией интерфейса GistLister
type MockGistLister struct{}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	gists := []*github.Gist{
		{
			Description: github.String("Gist1"),
			HTMLURL:     github.String("http://gist.github.com/user/gist1"),
		},
		{
			Description: github.String("Gist2"),
			HTMLURL:     github.String("http://gist.github.com/user/gist2"),
		},
		{
			Description: github.String("Gist3"),
			HTMLURL:     github.String("http://gist.github.com/user/gist3"),
		},
		{
			Description: github.String("Gist4"),
			HTMLURL:     github.String("http://gist.github.com/user/gist4"),
		},
	}
	return gists, nil, nil
}

func TestGithubAdapter_GetGists(t *testing.T) {
	type fields struct {
		RepoList RepoLister
		GistList GistLister
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				RepoList: &MockRepoLister{},
				GistList: &MockGistLister{},
			},
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			want: []Item{
				{
					Title:       "Gist1",
					Description: "Gist1",
					Link:        "http://gist.github.com/user/gist1",
				},
				{
					Title:       "Gist2",
					Description: "Gist2",
					Link:        "http://gist.github.com/user/gist2",
				},
				{
					Title:       "Gist3",
					Description: "Gist3",
					Link:        "http://gist.github.com/user/gist3",
				},
				{
					Title:       "Gist4",
					Description: "Gist4",
					Link:        "http://gist.github.com/user/gist4",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubAdapter{
				RepoList: tt.fields.RepoList,
				GistList: tt.fields.GistList,
			}
			got, err := g.GetGists(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubAdapter_GetRepos(t *testing.T) {
	type fields struct {
		RepoList RepoLister
		GistList GistLister
	}
	type args struct {
		ctx      context.Context
		username string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				RepoList: &MockRepoLister{},
				GistList: &MockGistLister{},
			},
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			want: []Item{
				{
					Title:       "Repo1",
					Description: "Description1",
					Link:        "http://github.com/user/repo1",
				},
				{
					Title:       "Repo2",
					Description: "Description2",
					Link:        "http://github.com/user/repo2",
				},
				{
					Title:       "Repo3",
					Description: "Description3",
					Link:        "http://github.com/user/repo3",
				},
				{
					Title:       "Repo4",
					Description: "Description4",
					Link:        "http://github.com/user/repo4",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubAdapter{
				RepoList: tt.fields.RepoList,
				GistList: tt.fields.GistList,
			}
			got, err := g.GetRepos(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubProxy_GetGists(t *testing.T) {
	type fields struct {
		github Githuber
		cache  map[string][]Item
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				github: &GithubAdapter{
					RepoList: &MockRepoLister{},
					GistList: &MockGistLister{},
				},
				cache: make(map[string][]Item),
			},
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			want: []Item{
				{
					Title:       "Gist1",
					Description: "Gist1",
					Link:        "http://gist.github.com/user/gist1",
				},
				{
					Title:       "Gist2",
					Description: "Gist2",
					Link:        "http://gist.github.com/user/gist2",
				},
				{
					Title:       "Gist3",
					Description: "Gist3",
					Link:        "http://gist.github.com/user/gist3",
				},
				{
					Title:       "Gist4",
					Description: "Gist4",
					Link:        "http://gist.github.com/user/gist4",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubProxy{
				github: tt.fields.github,
				cache:  tt.fields.cache,
			}
			got, err := g.GetGists(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubProxy_GetRepos(t *testing.T) {
	type fields struct {
		github Githuber
		cache  map[string][]Item
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				github: &GithubAdapter{
					RepoList: &MockRepoLister{},
					GistList: &MockGistLister{},
				},
				cache: make(map[string][]Item),
			},
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			want: []Item{
				{
					Title:       "Repo1",
					Description: "Description1",
					Link:        "http://github.com/user/repo1",
				},
				{
					Title:       "Repo2",
					Description: "Description2",
					Link:        "http://github.com/user/repo2",
				},
				{
					Title:       "Repo3",
					Description: "Description3",
					Link:        "http://github.com/user/repo3",
				},
				{
					Title:       "Repo4",
					Description: "Description4",
					Link:        "http://github.com/user/repo4",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubProxy{
				github: tt.fields.github,
				cache:  tt.fields.cache,
			}
			got, err := g.GetRepos(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGithubProxy(t *testing.T) {
	type args struct {
		github Githuber
	}
	tests := []struct {
		name string
		args args
		want *GithubProxy
	}{
		{
			name: "test1",
			args: args{
				github: &GithubAdapter{
					RepoList: &MockRepoLister{},
					GistList: &MockGistLister{},
				},
			},
			want: &GithubProxy{
				github: &GithubAdapter{
					RepoList: &MockRepoLister{},
					GistList: &MockGistLister{},
				},
				cache: make(map[string][]Item),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGithubProxy(tt.args.github); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGithubProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
