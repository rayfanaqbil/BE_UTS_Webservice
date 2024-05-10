package _gamesdatanih

import (
	"fmt"
	"testing"
)

func TestInsertDataGame(t *testing.T) {
	title := "Hitman Blood Money"
	genre := "Action 	"
	developer := "IO Interactive"
	publisher := "IO Interactive"
	releaseyear := 2006
	platform := "PC"
	mode := "Single Player"
	price := 45.99
	rating := 8.5

	insertedID := InsertDataGame(title, genre, developer, publisher, releaseyear, platform, mode, price, rating)
	fmt.Println(insertedID)
}

func TestGetGamesByTitle(t *testing.T) {
    title := "Hitman Blood Money"
    game := GetDataByTitle(title)
    fmt.Println(game)
}

func TestGetAll(t *testing.T) {
	data := GetAllDataGames()
	fmt.Println(data)
}
