package main

import (
	"context"
	"fmt"
	"time"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	},
	)

	if err != nil {
		return fmt.Errorf("Experienced error while to save feed to database: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	},
	)

	if err != nil {
		return fmt.Errorf("Experienced error while to saving feed follow to database: %w", err)
	}

	return nil
}
