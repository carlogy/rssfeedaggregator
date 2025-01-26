package main

import "fmt"

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.cmdMap[name] = f

}

func (c *commands) run(s *state, cmd command) error {

	if len(cmd.name) == 0 {
		return fmt.Errorf("command name %s is empty, pass valid command", cmd.name)
	}

	err := c.cmdMap[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}
