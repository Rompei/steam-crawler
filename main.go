package main

import (
	"github.com/Rompei/steam-crawler/crawler"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := crawler.NewCrawler()
	err := c.StartCrawl()
	if err != nil {
		panic(err)
	}
	err = c.StoreCSV("data.csv")
	if err != nil {
		panic(err)
	}
}
