package database

import (
	"back/src/hashing"
	"context"
	"errors"
	"fmt"
	"log"

	_ "back/src/hashing"
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

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Login    string             `json:"login"`
	Password string             `json:"password"`
}

var ctx = context.TODO()
var opts = options.Client().ApplyURI(uri)

func HealthCheck() {
	// Инициализируем клиент для БД
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	// Пингуем базу для проверки соединения
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	// Разрываем соединение c БД
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Successfully connected and pinged")
}

func GetGames(gameId string) []Game {
	var result []Game
	var filter bson.M
	// Инициализируем клиент
	client, _ := mongo.Connect(ctx, opts)
	// Подключаемся к базе games с коллекцией games_collection
	col := client.Database("games").Collection("games_collection")
	// Если в кваери параметрах присутствует ID, то добавляем его в фильтр при поиске игры в БД
	if gameId != "" {
		objectId, err := primitive.ObjectIDFromHex(gameId)
		if err != nil {
			log.Println("Could not parse query param to ObjectID")
		}
		filter = bson.M{"_id": objectId}
	}
	// Ищем игры в базе
	cur, err := col.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	// Добавляем все найденные игры в слайс result
	for cur.Next(ctx) {
		var elem Game
		err := cur.Decode(&elem)
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

func AddNewGame(reqBody Game) (string, error) {
	// Инициализируем клиент
	client, _ := mongo.Connect(ctx, opts)
	// Подключаемся к базе games с коллекцией games_collection
	col := client.Database("games").Collection("games_collection")
	// Создаем в коллекции запись игры с переданными параметрами
	result, err := col.InsertOne(ctx, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	// Получаем ID созданной игры
	gameId := result.InsertedID.(primitive.ObjectID).Hex()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return gameId, err
}

func DeleteGame(reqBody Game) (int64, error) {
	// Инициализируем клиент
	client, _ := mongo.Connect(ctx, opts)
	// Подключаемся к базе games с коллекцией games_collection
	col := client.Database("games").Collection("games_collection")
	// Удаляем игру с указанным ID
	result, err := col.DeleteOne(ctx, bson.M{"_id": reqBody.ID})
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return result.DeletedCount, err
}

func RegisterNewUser(user User) (string, error) {
	var userID string
	var handledError error
	// Хэшируем пароль пользователя
	user.Password = hashing.HashPassword(user.Password)
	// Инициализируем клиент
	client, _ := mongo.Connect(ctx, opts)
	// Подключаемся к базе games с коллекцией users
	col := client.Database("games").Collection("users")
	// Создаем пользователя с указанными параметрами
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		if errors.As(mongo.WriteError{}, &err) {
			handledError = errors.New("duplicated login")
			return userID, handledError
		} else {
			log.Fatal(err)
		}
	}
	// Получаем ID созданного пользователя
	userID = result.InsertedID.(primitive.ObjectID).Hex()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return userID, handledError
}
