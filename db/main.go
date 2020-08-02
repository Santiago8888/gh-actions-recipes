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

	"github.com/joho/godotenv"
)

const prefix = "mongodb://Admin"
const dbName = "actions"

const shard_0 = "cluster0-shard-00-00-8pgr7.mongodb.net:27017,"
const shard_1 = "cluster0-shard-00-01-8pgr7.mongodb.net:27017,"
const shard_2 = "cluster0-shard-00-02-8pgr7.mongodb.net:27017"
const cluster = shard_0 + shard_1 + shard_2

const config = "ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
const collName = "metrics"

var collection *mongo.Collection

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
	fmt.Println(results)
}
