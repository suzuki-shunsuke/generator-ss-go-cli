package main

import (
	"os"

	"github.com/urfave/cli"

	"<%= package %>/cmds"
)

func main() {
	app := cli.NewApp()
	app.Usage = ""
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "create a configuration file",
			Action: cmds.Init,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "The path of the configuration file",
				},
				cli.BoolFlag{
					Name:  "dry-run",
					Usage: "",
				}},
		},
	}
	app.Run(os.Args)
}
