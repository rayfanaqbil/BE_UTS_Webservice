package _gamesdatanih

import (
	"fmt"
	"testing"
)

func TestInsertDataGame(t *testing.T) {
	title := "Resident Evil Remake 3"
	genre := "Horror"
	developer := "Capcom"
	publisher := "Capcom"
	releaseyear := 2022
	platform := "PC"
	mode := "Single Player"
	price := 75.99
	rating := 7.5

	insertedID := InsertDataGame(title, genre, developer, publisher, releaseyear, platform, mode, price, rating)
	fmt.Println(insertedID)
}

func TestGetGamesByTitle(t *testing.T) {
    title := "Resident Evil Remake 3"
    game := GetGamesByTitle(title)
    fmt.Println(game)
}
