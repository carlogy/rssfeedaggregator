package main

import (
	"fmt"
)

func argsScrubber(args []string) (string, []string, error) {

	if len(args) < 2 {
		return "", nil, fmt.Errorf("Invalid amount of argumentas passed in %v.\nPlease submit a valid command and any of it's required arguments along with the program name.", args)
	}

	cmd := args[1]
	cmdArgs := args[2:]

	return cmd, cmdArgs, nil

}
