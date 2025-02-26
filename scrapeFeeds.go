package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/carlogy/rssfeedaggregator/internal/database"

	rf "github.com/carlogy/rssfeedaggregator/internal/RSS"
)

func scrapeFeeds(s *state) {

	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            feedToFetch.ID,
	})

	if err != nil {
		fmt.Println(fmt.Errorf("Error marking feed fetched: %w", err))
	}

	feed, err := rf.FetchFeed(context.Background(), feedToFetch.Url)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range feed.Channel.Item {

		// fmt.Println(item.Title)

		var valid bool
		if item.Description != "" {
			valid = true
		}

		publishDate, _ := time.Parse("1/2/2006 15:04:05", item.PubDate)

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Link,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: valid},
			PublishedAt: publishDate,
			FeedID:      feedToFetch.ID,
		})

		if err != nil {
			continue
		}

	}

}
