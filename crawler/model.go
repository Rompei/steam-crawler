package crawler

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

// Game is object to store steam game
type Game struct {
	ID            bson.ObjectId `json:"name" bosn:"_id"`
	Name          string        `json:"name" bosn:"name"`
	ReleaseDate   string        `json:"releaseDate" bson:"release_date"`
	Number        int           `json:"number" "bson:"number"`
	DiscountRate  int           `json:"discountRate" bson:"discount_rate"`
	NormalPrice   int           `json:"normalPrice" bson:"normal_price"`
	DiscountPrice int           `json:"discountPrice" bson:"discount_price"`
	Rate          int           `json:"rate" bson:"rate"`
	Reviewer      int           `json:"reviewer" bson:"reviewer"`
}

func (g *Game) String() {
	fmt.Printf("\n")
	fmt.Printf("Title: %s\n", g.Name)
	fmt.Printf("Index: %d\n", g.Number)
	fmt.Printf("Release date: %s\n", g.ReleaseDate)
	fmt.Printf("Discount rate: %d%%\n", g.DiscountRate)
	fmt.Printf("Normal price: %dYEN\n", g.NormalPrice)
	fmt.Printf("Discount price: %dYEN\n", g.DiscountPrice)
	fmt.Printf("Reputation: %d/100\n", g.Rate)
	fmt.Printf("The number of reviews: %d\n", g.Reviewer)
	fmt.Printf("\n")
}
