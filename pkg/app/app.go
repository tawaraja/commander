package app

import "github.com/urfave/cli"

const (
	//AppName defines the app name
	AppName = "Commander"
	//CommanderFile holds the default config file which is loaded
	CommanderFile = "commander.yaml"
)

//AddCommandContext holds all flags for the add command
type AddCommandContext struct {
	Verbose    bool
	NoColor    bool
	Concurrent int
}

//NewAddContextFromCli is a constructor which creates the context
func NewAddContextFromCli(c *cli.Context) AddCommandContext {
	return AddCommandContext{
		Verbose:    c.Bool("verbose"),
		NoColor:    c.Bool("no-color"),
		Concurrent: c.Int("concurrent"),
	}
}
