package _gamesdatanih


import(
	"fmt"
	"testing"
)

func TestInsertDataGame(t *testing.T) {
    title := "Resident Evil Remake 2"
    genre := "Horror"
    developer := "Capcom"
    publisher := "Capcom"
    releaseyear := 2019
    platform := "PC"
    mode := "Single Player"
    price := 49.99
    rating := 5.5

    insertedID := InsertDataGame(title, genre, developer, publisher, releaseyear, platform, mode, price, rating)
    fmt.Println(insertedID)
}

func TestGetAll(t *testing.T) {
	data := GetAllGames()
	fmt.Println(data)
}