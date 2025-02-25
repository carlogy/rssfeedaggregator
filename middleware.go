package main

import (
	"context"
	"fmt"

	"github.com/carlogy/rssfeedaggregator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User, ) error) (func (s* state, cmd command) error) {


		return func (s *state, cmd command) error {
			user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

			if err != nil {
				return fmt.Errorf("Error getting user: %w", err)
			}

			return handler(s, cmd, user)
		}
}
