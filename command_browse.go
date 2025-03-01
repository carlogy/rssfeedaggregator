package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int
	if len(cmd.arguments) < 1 {
		limit = 2
	} else {

		limit, _ = strconv.Atoi(cmd.arguments[0])
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))

	if err != nil {
		return fmt.Errorf("Experienced error: %w while browsing for posts", err)
	}

	for _, post := range posts {
		fmt.Printf("\nPost ID:\t%d\nPost CreatedAt:\t%v\nPost UpdatedAt:\t%v\nPost Title:\t%s\nPost Description:\t%s\nPublished:\t%v\nFrom Feed:\t%d\n\n", post.ID, post.CreatedAt, post.UpdatedAt, post.Title, post.Description.String, post.PublishedAt, post.FeedID)
	}

	return nil
}
