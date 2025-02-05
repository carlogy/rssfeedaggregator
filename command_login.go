package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogins(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("Error: login command requires username. No username argument was passed with the command.")

	}

	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])

	if err != nil {
		return errors.New("Error: user name provided, doesn't exist. Please submit a valid user")
	}

	s.config.SetUser(user.Name)

	fmt.Printf("%s was set as user\n", user.Name)

	return nil
}
