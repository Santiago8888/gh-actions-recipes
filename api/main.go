package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

type PrStats struct {
	// Repo Id
	owner      string
	repository string

	// PR Metadata
	number                int
	state                 string
	merged                bool
	title                 string
	created_at            time.Time
	closed_at             time.Time
	author_association    string
	maintainer_can_modify bool

	// PR Stats
	assignees_count           int
	requested_reviewers_count int
	comments                  int
	review_comments           int
	commits                   int
	additions                 int
	deletions                 int
	changed_files             int

	// Computed Stats
	time_diff  float64
	lines_diff int
}

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	owner := "cypress-io"
	repo := "cypress"

	opt := &github.PullRequestListOptions{State: "closed"}
	prs, _, _ := client.PullRequests.List(ctx, owner, repo, opt)

	pr := prs[0]
	prStat := &PrStats{}

	prStat.repository = repo
	prStat.owner = owner

	prStat.number = pr.GetNumber()
	prStat.state = pr.GetState()
	prStat.merged = pr.GetMerged()
	prStat.title = pr.GetTitle()
	prStat.created_at = pr.GetCreatedAt()
	prStat.closed_at = pr.GetClosedAt()
	prStat.author_association = pr.GetAuthorAssociation()
	prStat.maintainer_can_modify = pr.GetMaintainerCanModify()

	prStat.assignees_count = len(pr.Assignees)
	prStat.requested_reviewers_count = len(pr.RequestedReviewers)
	prStat.comments = pr.GetComments()
	prStat.review_comments = pr.GetReviewComments()
	prStat.commits = pr.GetCommits()
	prStat.additions = pr.GetAdditions()
	prStat.deletions = pr.GetDeletions()

	prStat.time_diff = prStat.closed_at.Sub(prStat.created_at).Hours()
	prStat.lines_diff = prStat.additions - prStat.deletions

	fmt.Println(prStat)
}
