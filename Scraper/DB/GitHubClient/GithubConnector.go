package GitHubClient

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/ancalabrese/Gypsy-19/Scraper/Data"
	"github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"
	"github.com/google/go-github/v37/github"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type GitHubConnector struct {
	logger     hclog.Logger
	client     *github.Client
	repo       *Repo
	commitInfo *CommitInfo
	ctx        context.Context
}

type Repo struct {
	Name       string `yaml:"name"`
	Owner      string `yaml:"owner"`
	BaseBranch string `yaml:"base-branch"`
	FilePath   string `yaml:"db-file-path"`
}

type CommitInfo struct {
	ClientName      string `yaml:"name"`  //Name for this client used for commiter name
	ClientEmail     string `yaml:"email"` //Email address used for commiter email addres
	commitMessage   string //Standard messsage for new commit 12 July 2021 - Travel list update
	CommitingBranch string `yaml:"branch"` //The branch where we are going to commit to
}

func NewGitHubConnector(l hclog.Logger, committer *CommitInfo, repo *Repo) *GitHubConnector {
	ctx := context.Background()
	ghToken, _ := viper.Get("GITHUB_TOKEN").(string)
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	httpClient := oauth2.NewClient(ctx, ts)

	committer.commitMessage = time.Now().Format("January 2, 2006") + " - Travel list update"
	ghc := &GitHubConnector{
		logger:     l,
		client:     github.NewClient(httpClient),
		ctx:        ctx,
		commitInfo: committer,
		repo:       repo,
	}

	ghc.logger.Debug("Created a new GitHub connector")
	return ghc
}

func (ghc *GitHubConnector) UpdateDB(lists Country.Lists) error {
	ghc.logger.Info("Updating country lists")
	ref, _ := ghc.getRef()
	tree, _ := ghc.getTree(ref)
	if err := ghc.updateLists(lists, tree); err != nil {
		return err
	}
	ghc.logger.Info("Finished updating repo")
	return nil
}

//getRef returns the commit branch reference or it creates it
//starting from the base branch
func (ghc *GitHubConnector) getRef() (ref *github.Reference, err error) {
	ghc.logger.Debug("Retrieving commit branch reference...", "Repo", ghc.repo.Name, "Ref", ghc.commitInfo.CommitingBranch)
	if ref, _, err = ghc.client.Git.GetRef(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, "heads/"+ghc.commitInfo.CommitingBranch); err == nil {
		return ref, nil
	}
	// An error means we couldnt retrieve the speicified ref because it does not exitsts. Creating a new ref
	ghc.logger.Debug("Failed retrieving reference. Creating new branch", "Repo", ghc.repo.Name, "Ref", ghc.commitInfo.CommitingBranch)
	var baseRef *github.Reference

	if baseRef, _, err = ghc.client.Git.GetRef(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, "/refs/heads"+ghc.repo.BaseBranch); err != nil {
		ghc.logger.Debug("Failed retrieving reference.", "Repo", ghc.repo.Name, "Ref", ghc.repo.BaseBranch)
		return nil, err
	}
	newRef := &github.Reference{Ref: github.String("refs/heads/" + ghc.commitInfo.CommitingBranch), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = ghc.client.Git.CreateRef(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, newRef)
	if err != nil {
		ghc.logger.Debug("Failed creating new reference.", "Repo", ghc.repo.Name, "Ref", ghc.commitInfo.CommitingBranch)
		return nil, err
	}
	return ref, err
}

// getTree gets the GitHub tree for the given reference returned from getRef()
func (ghc *GitHubConnector) getTree(ref *github.Reference) (tree *github.Tree, err error) {
	ghc.logger.Debug("Getting Tree", "Ref", ref.URL)
	tree, _, err = ghc.client.Git.GetTree(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, *ref.Object.SHA, false)
	return tree, err
}

func (ghc *GitHubConnector) downloadContent(filePath string) (*Country.Lists, error) {
	data, _, err := ghc.client.Repositories.DownloadContents(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, filePath, nil)
	if err != nil {
		ghc.logger.Error("Failed to download content", "FilePath", filePath, "Repo", ghc.repo.Name)
		return nil, err
	}
	defer data.Close()
	buf := new(strings.Builder)
	_, err = io.Copy(buf, data)
	// check errors
	ghc.logger.Debug("downloaded", "file", buf.String())

	lists := &Country.Lists{}
	err = Data.FromJson(lists, data)
	ghc.logger.Debug("Lists", "List", lists)
	if err != nil {
		ghc.logger.Error("Unable to unmarshal data", "FilePath", filePath, "Repo", ghc.repo.Name, "Error", err)
		return nil, err
	}
	return lists, nil
}

func (ghc *GitHubConnector) updateLists(data Country.Lists, tree *github.Tree) error {
	var sb strings.Builder
	var SHA string
	for i := range tree.Entries {
		if tree.Entries[i].GetPath() == ghc.repo.FilePath {
			SHA = tree.Entries[i].GetSHA()
			break
		}
	}
	Data.ToPrettyJson(data, &sb)
	fileOptions := &github.RepositoryContentFileOptions{
		Message:   &ghc.commitInfo.commitMessage,
		Content:   []byte(sb.String()),
		SHA:       &SHA,
		Branch:    &ghc.commitInfo.CommitingBranch,
		Author:    &github.CommitAuthor{Name: &ghc.commitInfo.ClientName, Email: &ghc.commitInfo.ClientEmail},
		Committer: &github.CommitAuthor{Name: &ghc.commitInfo.ClientName, Email: &ghc.commitInfo.ClientEmail},
	}
	r, s, err := ghc.client.Repositories.UpdateFile(ghc.ctx, ghc.repo.Owner, ghc.repo.Name, ghc.repo.FilePath, fileOptions)
	if err != nil {
		ghc.logger.Error("Unable to update file", "FilePath", tree.Entries[0].GetPath(), "Repo", ghc.repo.Name, "Error", err, r, s)
		return err
	}
	ghc.logger.Info("Updates pushed.")
	return nil
}
