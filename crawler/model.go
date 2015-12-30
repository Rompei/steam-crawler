package crawler

import (
	"fmt"
	"strconv"
)

// Game is object to store steam game
type Game struct {
	Number        int    `json:"number"`
	Name          string `json:"name"`
	ReleaseDate   string `json:"releaseDate"`
	DiscountRate  int    `json:"discountRate"`
	NormalPrice   int    `json:"normalPrice"`
	DiscountPrice int    `json:"discountPrice"`
	Rate          int    `json:"rate"`
	Reviewer      int    `json:"reviewer"`
	URL           string `json:"url"`
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
	fmt.Printf("URL: %s\n", g.URL)
	fmt.Printf("\n")
}

// GetRow returns the array of the game.
func (g *Game) GetRow() (row []string) {
	row = make([]string, 9)
	row[0] = strconv.Itoa(g.Number)
	row[1] = g.Name
	row[2] = g.ReleaseDate
	row[3] = strconv.Itoa(g.DiscountRate)
	row[4] = strconv.Itoa(g.NormalPrice)
	row[5] = strconv.Itoa(g.DiscountPrice)
	row[6] = strconv.Itoa(g.Rate)
	row[7] = strconv.Itoa(g.Reviewer)
	row[8] = g.URL
	return
}
