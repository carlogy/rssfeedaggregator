package main

import (
	"context"
	"fmt"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		fmt.Errorf("Experienced error: %w while getting feeds the user is following", err)
	}

	for _, follow := range following {
		fmt.Println(follow)
	}

	return nil
}
