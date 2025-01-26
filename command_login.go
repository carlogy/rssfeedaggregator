package main

import (
	"errors"
	"fmt"
)

func handlerLogins(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("Error: login command requires username. No username argument was passed with the command.")

	}

	s.config.SetUser(cmd.arguments[0])

	fmt.Printf("%s was set as user\n", cmd.arguments[0])

	return nil
}
