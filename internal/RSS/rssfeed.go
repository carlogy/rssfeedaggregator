package rss

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func NewRSSFeed() RSSFeed {
	return RSSFeed{}
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")
	req.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	// fmt.Println(string(body))

	rssfeed := NewRSSFeed()
	xml.Unmarshal(body, &rssfeed)

	decodedRSSFeed, err := unescapeStrings(&rssfeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	return decodedRSSFeed, nil
}

func unescapeStrings(rf *RSSFeed) (*RSSFeed, error) {

	if rf == nil {
		return &RSSFeed{}, errors.New("Unable to decode escaped html entities, received nil instead reference to RSSFeed.")
	}

	rf.Channel.Title = html.UnescapeString(rf.Channel.Title)
	rf.Channel.Description = html.UnescapeString(rf.Channel.Description)

	for i := 0; i < len(rf.Channel.Item); i++ {
		title := html.UnescapeString(rf.Channel.Item[i].Title)
		description := html.UnescapeString(rf.Channel.Item[i].Description)

		rf.Channel.Item[i].Title = title
		rf.Channel.Item[i].Description = description
	}

	return rf, nil
}
