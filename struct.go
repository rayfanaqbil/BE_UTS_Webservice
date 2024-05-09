package _gamesdatanih


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameData struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Title       string             `bson:"title" json:"title"` 
    Genre       string             `bson:"genre" json:"genre"`
    Developer   string             `bson:"developer" json:"developer"`
    Publisher   string             `bson:"publisher" json:"publisher"`
    ReleaseYear int                `bson:"release_year" json:"release_year"`
    Platform    string             `bson:"platform" json:"platform"`
    Mode        string             `bson:"mode" json:"mode"`
    Price       float64            `bson:"price" json:"price"`
    Rating      float64            `bson:"rating" json:"rating"`
}