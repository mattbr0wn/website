package rss

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/mattbr0wn/website/config"
	"github.com/mattbr0wn/website/internal/markdown"
	"github.com/mattbr0wn/website/internal/ssg"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate string `xml:"lastBuildDate"`
	Items         []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func WriteRssFeed(markdownFiles *[]string) {
	xmlData, err := buildRssFeed(markdownFiles)
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(config.ROOT_DIR, "feed.xml")
	writeErr := os.WriteFile(outputPath, xmlData, 0644)
	if writeErr != nil {
		panic(writeErr)
	}
}

func getFeedItems(markdownFiles *[]string) []Item {
	var rssItems []Item

	for _, file := range *markdownFiles {
		metadata, _, _ := markdown.ParseMarkdownFile(file)
		if metadata.Draft != true && filepath.Base(file) != "index.md" {
			link := filepath.Join(config.WEBSITE_URL, strings.TrimPrefix(ssg.GenerateStaticUrl(file), config.ROOT_DIR))
			feedItems := Item{
				Title:       metadata.Title,
				Link:        link,
				Description: metadata.Description,
				PubDate:     convertDateToRFC1123Z(metadata.Date),
			}
			rssItems = append(rssItems, feedItems)
		}
	}

	return sortItemsByPubDate(rssItems)
}

func getChannelInfo(items []Item) Channel {
	lastDate := items[0].PubDate
	channel := Channel{
		Title:         config.TITLE,
		Link:          config.WEBSITE_URL,
		Description:   config.DESCRIPTION,
		Language:      "en-us",
		LastBuildDate: lastDate,
		Items:         items,
	}

	return channel
}

func sortItemsByPubDate(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool {
		dateI, _ := time.Parse(time.RFC1123Z, items[i].PubDate)
		dateJ, _ := time.Parse(time.RFC1123Z, items[j].PubDate)
		return dateI.After(dateJ)
	})
	return items
}

func convertDateToRFC1123Z(dateString string) string {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Printf("ERROR: Cannot parse date %v", err)
		return ""
	}
	return date.Format(time.RFC1123Z)
}

func buildRssFeed(markdownFiles *[]string) ([]byte, error) {
	rssItems := getFeedItems(markdownFiles)
	channel := getChannelInfo(rssItems)

	rss := RSS{
		Version: "2.0",
		Channel: channel,
	}

	// Marshal the RSS struct to XML
	xmlData, err := xml.Marshal(rss)

	return xmlData, err
}
