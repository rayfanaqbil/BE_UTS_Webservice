package _gamesdatanih

import (
	"context"
	"fmt"
	"os"

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

func InsertDataGame(title string,genre string,developer string,publisher string,realeseyear int, platform string,mode string,price float64,rating float64) (InsertedID interface{}) {
	var gamedata GameData
	gamedata.Title = title
	gamedata.Genre = genre
	gamedata.Developer = developer
	gamedata.Publisher = publisher
	gamedata.ReleaseYear = realeseyear
	gamedata.Platform = platform
	gamedata.Mode = mode
	gamedata.Price = price
	gamedata.Rating = rating
	return InsertOneDoc("DataGame", "Game", gamedata )
}

func GetAllGames() (data []GameData) {
	gem := MongoConnect("DataGame").Collection("Game")
	filter := bson.M{}
	cursor, err := gem.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetGamebyID(id primitive.ObjectID) (GameData, error) {
    var game GameData

    filter := bson.M{"_id": id}
    err := MongoConnect("DataGame").Collection("Game").FindOne(context.Background(), filter).Decode(&game)
    if err != nil {
        return GameData{}, err
    }

    return game, nil
}