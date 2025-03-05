package main

import (
	"database/sql"
	"fmt"
	"os"

	c "github.com/carlogy/rssfeedaggregator/internal/config"
	database "github.com/carlogy/rssfeedaggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := c.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)

	if err != nil {
		fmt.Println("Experienced error connecting to db:\t", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := NewState(&cfg, dbQueries)

	commands := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogins)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

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

}
