package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	if cmd.name != "users" {
		return errors.New("Error: Invalid command passed.")
	}

	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return errors.New("Error: failed to get users list from users table")
	}

	for _, user := range users {

		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		}

		fmt.Printf("* %s", user.Name)
	}

	return nil

}
