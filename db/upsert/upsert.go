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
	Owner      string `bson:"owner"`
	Repository string `bson:"repository"`

	// PR Metadata
	Number              int       `bson:"number"`
	State               string    `bson:"state"`
	Merged              bool      `bson:"merged"`
	Title               string    `bson:"title"`
	CreatedAt           time.Time `bson:"created_at"`
	ClosedAt            time.Time `bson:"closed_at"`
	AuthorAssociation   string    `bson:"author_association"`
	MaintainerCanModify bool      `bson:"maintainer_can_modify"`

	// PR Stats
	AssigneesCount          int `bson:"assignees_count"`
	RequestedReviewersCount int `bson:"requested_reviewers_count"`
	Comments                int `bson:"comments"`
	ReviewComments          int `bson:"review_comments"`
	Commits                 int `bson:"commits"`
	Additions               int `bson:"additions"`
	Deletions               int `bson:"deletions"`
	ChangedFiles            int `bson:"changed_files"`

	// Computed Stats
	TimeDiff  float64 `bson:"time_diff"`
	LinesDiff int     `bson:"lines_diff"`
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

	repo := "cypress"
	cursor, err := collection.Find(context.TODO(), bson.D{primitive.E{Key: "repository", Value: repo}})
	if err != nil {
		log.Fatal(err)
	}

	var results []PrStat
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	gitCtx := context.Background()
	client := github.NewClient(nil)

	for _, result := range results {
		pr, _, _ := client.PullRequests.Get(gitCtx, result.Owner, result.Repository, result.Number)

		var prDoc PrStat
		prDoc.Repository = result.Owner
		prDoc.Owner = result.Repository

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

		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.D{primitive.E{Key: "repository", Value: repo}, primitive.E{Key: "number", Value: pr.Number}}
		update := bson.D{primitive.E{Key: "$set", Value: prDoc}}

		collection.FindOneAndUpdate(context.TODO(), filter, update, opts)
	}
}
