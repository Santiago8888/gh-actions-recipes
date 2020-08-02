package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
)

const prefix = "mongodb://Admin"
const dbName = "actions"

const shard_0 = "cluster0-shard-00-00-8pgr7.mongodb.net:27017,"
const shard_1 = "cluster0-shard-00-01-8pgr7.mongodb.net:27017,"
const shard_2 = "cluster0-shard-00-02-8pgr7.mongodb.net:27017"
const cluster = shard_0 + shard_1 + shard_2

const config = "ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
const collName = "stats"

var collection *mongo.Collection

type PrStat struct {
	// Repo Id
	Owner      string `json:"owner,omitempty"`
	Repository string `json:"repository,omitempty"`

	// PR Metadata
	Number              int       `json:"number"`
	State               string    `json:"state"`
	Merged              bool      `json:"merged"`
	Title               string    `json:"title"`
	CreatedAt           time.Time `json:"created_at"`
	ClosedAt            time.Time `json:"closed_at"`
	AuthorAssociation   string    `json:"author_association"`
	MaintainerCanModify bool      `json:"maintainer_can_modify"`

	// PR Stats
	AssigneesCount          int `json:"assignees_count"`
	RequestedReviewersCount int `json:"requested_reviewers_count"`
	Comments                int `json:"comments"`
	ReviewComments          int `json:"review_comments"`
	Commits                 int `json:"commits"`
	Additions               int `json:"additions"`
	Deletions               int `json:"deletions"`
	ChangedFiles            int `json:"changed_files"`

	// Computed Stats
	TimeDiff  float64 `json:"time_diff"`
	LinesDiff int     `json:"lines_diff"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pwd := os.Getenv("MONGO_PWD")
	fmt.Println(pwd)
	connectionString := prefix + ":" + pwd + "@" + cluster + "/" + dbName + "?" + config

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = db.Database(dbName).Collection(collName)

	gitCtx := context.Background()
	client := github.NewClient(nil)

	owner := "ansible"
	repo := "ansible"

	opt := &github.PullRequestListOptions{State: "closed"}
	prs, _, _ := client.PullRequests.List(gitCtx, owner, repo, opt)

	for _, pr := range prs {
		var prDoc PrStat

		prDoc.Repository = repo
		prDoc.Owner = owner

		prDoc.Number = pr.GetNumber()
		prDoc.State = pr.GetState()
		prDoc.Merged = pr.GetMerged()
		prDoc.Title = pr.GetTitle()
		prDoc.CreatedAt = pr.GetCreatedAt()
		prDoc.ClosedAt = pr.GetClosedAt()
		prDoc.AuthorAssociation = pr.GetAuthorAssociation()
		prDoc.MaintainerCanModify = pr.GetMaintainerCanModify()

		prDoc.AssigneesCount = len(pr.Assignees)
		prDoc.RequestedReviewersCount = len(pr.RequestedReviewers)
		prDoc.Comments = pr.GetComments()
		prDoc.ReviewComments = pr.GetReviewComments()
		prDoc.Commits = pr.GetCommits()
		prDoc.Additions = pr.GetAdditions()
		prDoc.Deletions = pr.GetDeletions()

		prDoc.TimeDiff = prDoc.ClosedAt.Sub(prDoc.CreatedAt).Hours()
		prDoc.LinesDiff = prDoc.Additions - prDoc.Deletions

		insertPR(prDoc)
	}

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	fmt.Println(len(results))
}

func insertPR(pr PrStat) {
	insertResult, err := collection.InsertOne(context.Background(), pr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
