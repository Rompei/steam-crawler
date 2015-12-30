package crawler

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// URL for getting special games from steam.
var URL = "http://store.steampowered.com/search/results?sort_by=_ASC&specials=1"

// Crawler is steam crawler.
type Crawler struct {
	games []Game
}

// NewCrawler is constructor of Crawler.
func NewCrawler() *Crawler {
	return &Crawler{}
}

// GetGames return got games.
func (c *Crawler) GetGames() []Game {
	return c.games
}

// ShowAllGames shows all games.
func (c *Crawler) ShowAllGames() {
	for _, g := range c.games {
		g.String()
	}
}

// StartCrawl start crawling.
func (c *Crawler) StartCrawl() (err error) {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		return
	}

	// Getting the number of pages.
	var pageNum int
	doc.Find(".search_pagination_right").Children().Each(func(i int, s *goquery.Selection) {
		if i == 2 {
			pageNum, err = strconv.Atoi(s.Text())
			if err != nil {
				return
			}
		}
	})

	// Starting goroutine.
	resultCh := make(chan []Game, pageNum)
	for i := 1; i < pageNum+1; i++ {
		url := fmt.Sprintf("%s&page=%d", URL, i)
		go c.crawl(url, resultCh)
	}

	// Collectiong games.
	for i := 1; i < pageNum+1; i++ {
		gs := <-resultCh
		c.games = append(c.games, gs...)
	}
	close(resultCh)

	return
}

// StoreCSV stores game info as CSV.
func (c *Crawler) StoreCSV(path string) (err error) {
	c.SortGames()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	err = file.Truncate(0)
	if err != nil {
		return
	}
	writer := csv.NewWriter(file)
	writer.UseCRLF = true

	for _, v := range c.games {
		writer.Write(v.GetRow())
	}

	writer.Flush()
	return
}

// SortGames sorts games.
func (c *Crawler) SortGames() {
	c.games = sortData(c.games)
}

func sortData(games []Game) (ret []Game) {
	if len(games) == 0 {
		return games
	}

	pivot := games[0]

	var left []Game
	var right []Game

	for _, v := range games[1:] {
		if v.Number > pivot.Number {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	left = sortData(left)
	right = sortData(right)
	ret = append(left, pivot)
	ret = append(ret, right...)
	return
}

func (c *Crawler) crawl(url string, resultCh chan []Game) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	// Getting the number of element of a page
	elementNum, err := c.getFirstElementNumber(doc.Find(".search_pagination_left").Text())
	if err != nil {
		panic(err)
	}

	// Scraing
	var games []Game
	doc.Find(".search_result_row").Each(func(_ int, s *goquery.Selection) {
		var game Game

		// Getting title.
		game.Name = s.Find(".title").Text()

		// Getting release date.
		game.ReleaseDate = s.Find(".search_released").Text()

		// Getting discount rate.
		// There is a title not to be discounted Ex) Dark souls 2.
		game.DiscountRate, _ = c.extractDiscount(s.Find(".search_discount").Find("span").Text())

		// Getting normal price and discount price.
		game.NormalPrice, game.DiscountPrice, err = c.extractPrices(s.Find(".search_price").Text())
		if err != nil {
			panic(err)
		}

		// Getting user review.
		if review, exist := s.Find(".search_review_summary").Attr("data-store-tooltip"); exist == true {
			game.Rate, game.Reviewer, err = c.extractReview(review)
		}

		// Setting index of the games.
		game.Number = elementNum
		elementNum++

		// Getting link of the game
		if bareURL, exist := s.Attr("href"); exist == true {
			game.URL, err = c.extractURL(bareURL)
		}

		games = append(games, game)
	})

	resultCh <- games
}

// getFirstElementNumber gets index of the page
func (*Crawler) getFirstElementNumber(paginationLeft string) (int, error) {
	re, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		return 0, err
	}

	pageStr := re.FindString(paginationLeft)

	// FIX: There is cases can't treat the number of pages.
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, nil
	}
	return page, nil
}

// extractDiscount gets rate of discount.
func (*Crawler) extractDiscount(discount string) (int, error) {
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		return 0, err
	}
	exDiscount := re.FindString(discount)
	return strconv.Atoi(exDiscount)
}

// extractPrices return normal price and discount price.
func (*Crawler) extractPrices(price string) (normalPrice, discountPrice int, err error) {
	priceData := strings.Fields(price)
	re1, err := regexp.Compile(`\xa5|,|\$|[a-zA-Z]+`)
	if err != nil {
		return
	}

	if len(priceData) >= 2 {
		normalPrice, err = strconv.Atoi(re1.ReplaceAllString(priceData[1], ""))
		if err != nil {
			return
		}
		if len(priceData) != 3 {
			return
		}

		re2, err := regexp.Compile(`\xa5|,|\$|[a-zA-Z]+`)
		if err != nil {
			return normalPrice, discountPrice, err
		}

		discountPrice, _ = strconv.Atoi(re2.ReplaceAllString(priceData[2], ""))
	}

	return
}

// extractReview return the rate of good review.
func (*Crawler) extractReview(review string) (rate, reviewer int, err error) {
	re, err := regexp.Compile(`[0-9,]+`)
	if err != nil {
		return
	}

	numbers := re.FindAllString(review, -1)

	rate, err = strconv.Atoi(numbers[0])
	if err != nil {
		return
	}

	re1, err := regexp.Compile(`,`)
	if err != nil {
		return
	}

	reviewer, err = strconv.Atoi(re1.ReplaceAllString(numbers[1], ""))
	if err != nil {
		return
	}

	return
}

func (*Crawler) extractURL(bareURL string) (string, error) {
	parsedURL, err := url.Parse(bareURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path, nil
}
