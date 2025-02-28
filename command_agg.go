package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) < 1 {
		return errors.New("Error: agg command requires a time duration to scrape feeds")
	}

	duration, err := time.ParseDuration(cmd.arguments[0])

	if err != nil {
		return fmt.Errorf("Experienced Error %w, while parsing time duration", err)
	}
	fmt.Println("Collecting feeds every " + duration.String())
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		scrapeFeeds(s)
	}
	return nil
}
