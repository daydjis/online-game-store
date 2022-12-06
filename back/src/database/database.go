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
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	Price            int64              `json:"price"`
	Genres           []string           `json:"genres"`
	Image            string             `json:"image"`
	Video            string             `json:"video"`
	ImageDescription string             `json:"imageDescription"`
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

func GetGames(gameId string) []Game {
	var result []Game

	client, _ := mongo.Connect(ctx, opts)
	col := client.Database("games").Collection("games_collection")
	if gameId == "" {
		cur, err := col.Find(ctx, bson.D{{}})
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
	} else {
		var elem Game
		objectId, err := primitive.ObjectIDFromHex(gameId)
		if err != nil {
			log.Println("Could not parse query param to ObjectID")
		}
		err = col.FindOne(ctx, bson.M{"_id": objectId}).Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, elem)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
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
