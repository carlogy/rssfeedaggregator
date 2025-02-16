package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) < 1 {
		return errors.New("Error: url for feed was not provided along with unfollow command. Unable to remove follow")
	}
	url := cmd.arguments[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		return fmt.Errorf("Error while removing feed follow: %w", err)
	}

	return nil
}
