package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

const siteConfigPath = "config.yml"

type SiteConfig struct {
	Site struct {
		Url        string `yaml:"url"`
		DateFormat string `yaml:"date_format"`
	} `yaml:"site"`
	Selectors struct {
		EventRoot string `yaml:"event_root"`
		Title     string `yaml:"title"`
		DateAttr  string `yaml:"date_attr"`
	} `yaml:"selectors"`
}

type Event struct {
	Date  time.Time
	Title string
}

func loadSiteConfig(path string) (*SiteConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config SiteConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func loadDocument(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}
	return doc, nil
}

func fetchEvents(config *SiteConfig) ([]Event, error) {
	doc, err := loadDocument(config.Site.Url)
	if err != nil {
		return nil, err
	}

	var events []Event

	doc.Find(config.Selectors.EventRoot).Each(func(i int, s *goquery.Selection) {
		title := s.Find(config.Selectors.Title).Text()
		dateStr, exists := s.Attr(config.Selectors.DateAttr)
		if !exists || title == "" {
			return
		}
		date, err := time.Parse(config.Site.DateFormat, dateStr)
		if err != nil {
			log.Printf("Failed to parse date: %v", dateStr)
			return
		}
		events = append(events, Event{
			Title: title,
			Date:  date,
		})
	})

	// Sort events by date ascending
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})

	return events, nil
}

func main() {
	config, err := loadSiteConfig(siteConfigPath)
	if err != nil {
		log.Fatalf("Error loading site config: %v", err)
	}

	events, err := fetchEvents(config)
	if err != nil {
		log.Fatalf("Error fetching events: %v", err)
	}

	// Print events to check if they were correctly fetched and parsed
	if len(events) > 0 {
		for _, e := range events {
			fmt.Printf("%s | %s\n", e.Date.Format("2006-01-02"), e.Title)
		}
	} else {
		fmt.Println("No events found.")
	}
}
