package rss

import (
	"encoding/xml"
	"fmt"
	"log"
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
	xmlData, xmlErr := buildRssFeed(markdownFiles)
	if xmlErr != nil {
		log.Fatalf("Error building RSS feed: %v", xmlErr)
	}

	outputPath := filepath.Join(config.ROOT_DIR, "feed.xml")
	writeErr := os.WriteFile(outputPath, xmlData, 0644)
	if writeErr != nil {
		log.Fatalf("Error writing %s: %v", outputPath, writeErr)
	}
}

func getFeedItems(markdownFiles *[]string) []Item {
	var rssItems []Item

	for _, file := range *markdownFiles {
		parsedMarkdownFile, parseErr := markdown.ParseMarkdownFile(file)
		if parseErr != nil {
			log.Fatal(parseErr)
		}

		if parsedMarkdownFile.Frontmatter.Draft == false && filepath.Base(file) != config.INDEX {
			link := filepath.Join(config.WEBSITE_URL, strings.TrimPrefix(ssg.GenerateStaticPath(file), config.ROOT_DIR))
			pub_date, err := convertDateToRFC1123Z(parsedMarkdownFile.Frontmatter.Date)
			if err != nil {
				log.Println(err)
			}
			feedItems := Item{
				Title:       parsedMarkdownFile.Frontmatter.Title,
				Link:        link,
				Description: parsedMarkdownFile.Frontmatter.Description,
				PubDate:     pub_date,
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
		dateI, errI := time.Parse(time.RFC1123Z, items[i].PubDate)
		if errI != nil {
			log.Printf("Error parsing date %v: %v", items[i].PubDate, errI)
		}
		dateJ, errJ := time.Parse(time.RFC1123Z, items[j].PubDate)
		if errJ != nil {
			log.Printf("Error parsing date %v: %v", items[j].PubDate, errJ)
		}
		return dateI.After(dateJ)
	})
	return items
}

func convertDateToRFC1123Z(dateString string) (string, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return "", fmt.Errorf("Error parsing %v to string: %v", date, err)
	}
	return date.Format(time.RFC1123Z), nil
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
