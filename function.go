package _gamesdatanih

import (
	"context"
	"fmt"
	"os"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataGame(title string, genre string, developer string, publisher string, releaseyear int, platform string, mode string, price float64, rating float64) (InsertedID interface{}) {
	var gamedata GameData
	gamedata.Title = title
	gamedata.Genre = genre
	gamedata.Developer = developer
	gamedata.Publisher = publisher
	gamedata.ReleaseYear = releaseyear
	gamedata.Platform = platform
	gamedata.Mode = mode
	gamedata.Price = price
	gamedata.Rating = rating
	return InsertOneDoc("DataGame", "game", gamedata)
}

func GetAllDataGames() (data []GameData) {
	gem := MongoConnect("DataGame").Collection("game")
	filter := bson.M{}
	cursor, err := gem.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllData: ", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataByTitle(title string) (games GameData) {
	gem := MongoConnect("DataGame").Collection("game")
	filter := bson.M{"title": title}
	err := gem.FindOne(context.TODO(), filter).Decode(&games)
	if err != nil {
		fmt.Printf("getDataByTitle: %v\n", err)
	}
	return games
}

func GetGamesID(_id primitive.ObjectID, db *mongo.Database, col string) (games GameData, errs error) {
	gem := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := gem.FindOne(context.TODO(), filter).Decode(&games)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return games, fmt.Errorf("no data found for ID %s", _id)
		}
		return games, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return games, nil
}
