package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Experienced error: %w while attempting to query all feeds", err)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error: %w gettings users for feeds ", err)
	}

	usersMap := make(map[uuid.UUID]string)
	for _, user := range users {
		usersMap[user.ID] = user.Name
	}

	for _, feed := range feeds {
		userName := usersMap[feed.UserID]
		fmt.Printf("Feed Name: %s\tURL: %s\tAdded by user:\t%s\n", feed.Name, feed.Url, userName)
	}

	return nil
}
