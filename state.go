package main

import (
	cfg "github.com/carlogy/rssfeedaggregator/internal/config"
	"github.com/carlogy/rssfeedaggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *cfg.Config
}

func NewState(cfg *cfg.Config, db *database.Queries) state {
	return state{config: cfg, db: db}
}

type command struct {
	name      string
	arguments []string
}

func NewCommand(name string, arguments []string) command {
	return command{name: name, arguments: arguments}
}
