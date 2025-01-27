package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/carlogy/rssfeedaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: register command requires a user to register. No user argument was provided with register command.")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})

	if err != nil {
		return errors.New("Unable to register user, The user name is already registered")
	}

	s.config.SetUser(user.Name)

	userDetails := fmt.Sprintf("User\nUserID:\t%s\nUser Name:\t%s\nCreated:\t%vUpdated:\t%v", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	fmt.Println(userDetails)

	return nil
}
