package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	if cmd.name != "reset" {
		return fmt.Errorf("Error: invalid command passed in %s", cmd.name)
	}

	err := s.db.DeleteAllUsers(context.Background())

	if err != nil {
		return errors.New("Error: Experienced error while attempting to reset the users table.")
	}

	fmt.Println("Succesfully reset users table")
	return nil
}
