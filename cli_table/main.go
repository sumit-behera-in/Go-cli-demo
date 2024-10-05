package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pterm/pterm"
)

type News struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Articles struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

func main() {

	res, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=017688d9c55f46d493bfb4b3391d3c35")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var news News
	err = json.Unmarshal(bs, &news)

	var results []Articles = news.Articles

	if err != nil {
		log.Fatal(err)
	}

	tableData := pterm.TableData{
		{"Author", "Title", "Time"},
	}

	for _, article := range results {
		tableData = append(tableData, []string{
			func() string {
				if len(article.Author) > 15 {
					return article.Author[0:13] + "..." // Include 16th character for full 15 cuts
				}
				return article.Author
			}(),
			func() string {
				if len(article.Title) > 50 {
					return article.Title[0:47] + "..." // Include 16th character for full 15 cuts
				}
				return article.Title
			}(),
			fmt.Sprintf("%v-%v", article.PublishedAt.Month(), article.PublishedAt.Year()),
		})
	}

	err = pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData(tableData).
		Render()

	if err != nil {
		log.Fatal(err)
	}

}
