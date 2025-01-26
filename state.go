package main

import (
	cfg "github.com/carlogy/rssfeedaggregator/internal/config"
)

type state struct {
	config *cfg.Config
}

func NewState(cfg *cfg.Config) state {
	return state{config: cfg}
}

type command struct {
	name      string
	arguments []string
}

func NewCommand(name string, arguments []string) command {
	return command{name: name, arguments: arguments}
}
