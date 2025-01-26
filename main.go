package main

import (
	"fmt"
	"os"

	c "github.com/carlogy/rssfeedaggregator/internal/config"
)

func main() {

	cfg, err := c.Read()
	if err != nil {
		fmt.Println(err)
	}

	s := NewState(&cfg)

	commands := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogins)

	fmt.Println(s, commands)

	Args := os.Args

	cmd, cmdArgs, err := argsScrubber(Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command := NewCommand(cmd, cmdArgs)

	err = commands.run(&s, command)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// newConfigRead, error := c.Read()

	// if error != nil {
	// 	fmt.Println(error)
	// }

	// // fmt.Println(newConfigRead.CurrentUserName, newConfigRead.DbURL)

}
