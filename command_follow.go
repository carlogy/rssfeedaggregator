package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("Error: no url was supplied with the follow command.")
	}

	feeds, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])

	if err != nil {
		return fmt.Errorf("Experienced error: %w while attempting to get feed by provided url", err)
	}

	f, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feeds.ID,
	},
	)

	if err != nil {
		return fmt.Errorf("Experienced error: %w while attempting to follow feed", err)
	}

	fmt.Println(f.FeedName, f.UserName)

	return nil

}
