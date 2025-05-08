package event

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/leyka/basical/internal/config"
)

func fetchEvents(config *config.Config) ([]Event, error) {
	doc, err := loadDocument(config.Site.URL)
	if err != nil {
		return nil, err
	}

	selectors := config.Site.Selectors

	var events []Event
	doc.Find(selectors.EventRoot).Each(func(i int, s *goquery.Selection) {
		titleNode := s.Find(selectors.Title)
		title := titleNode.Text()
		link := titleNode.AttrOr("href", "")
		dateStr, dateExists := s.Attr(selectors.DateAttr)
		if !dateExists || title == "" {
			return
		}

		date, err := time.Parse(config.Site.DateFormat, dateStr)
		if err != nil {
			fmt.Printf("Failed to parse date: %s, error: %v\n", dateStr, err)
			return
		}
		events = append(events, Event{
			Title: title,
			Date:  date,
			Link:  link,
		})
	})

	// Sort events by date ascending
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})

	return events, nil
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
