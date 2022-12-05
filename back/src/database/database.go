package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://geniy_folomeev:26051978aA%21@gamestore.dsm2ow5.mongodb.net/?retryWrites=true&w=majority"

type Game struct {
	ID               primitive.ObjectID `bson:"_id"`
	Title            string
	Description      string
	Price            int64
	Genres           []string
	Image            string
	Video            string
	ImageDescription string
}

var ctx = context.TODO()
var opts = options.Client().ApplyURI(uri)

func HealthCheck() *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}

func GetGames() []Game {
	var result []Game
	findOptions := options.Find()

	client, _ := mongo.Connect(ctx, opts)
	col := client.Database("games").Collection("games_collection")
	cur, err := col.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		var elem Game
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, elem)

	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return result
}

func AddNewGame(v Game) {
	client, _ := mongo.Connect(ctx, opts)
	col := client.Database("games").Collection("games_collection")
	_, err := col.InsertOne(ctx, v)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}