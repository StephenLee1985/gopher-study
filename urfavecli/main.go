package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var greetCommand = cli.Command{
	Name:  "greet",
	Usage: "greet somebody",
	Action: func(c *cli.Context) error {
		name := "stupid egg"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Printf("Hola %s\n", name)

		} else {
			fmt.Printf("Hello, %s\n", name)

		}
		return nil
	},
}

var teachCommand = cli.Command{
	Name:  "teach",
	Usage: "teach somebody",
	Action: func(c *cli.Context) error {
		name := "stupid egg"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Printf("enseñar %s\n", name)

		} else {
			fmt.Printf("teach, %s\n", name)

		}
		return nil
	},
}

func main() {

	app := cli.NewApp()
	//.Run(os.Args)

	app.Name = "live house"
	app.Usage = "make a live house"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang,l",
			Value:  "english",
			Usage:  "language for greet",
			EnvVar: "APP_LANG,LANG",
		},
	}

	app.Commands = []cli.Command{
		greetCommand,
		teachCommand,
	}

	app.Run(os.Args)
	//fmt.Println("house morning")

}
