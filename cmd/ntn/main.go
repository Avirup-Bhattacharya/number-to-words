package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/moul/ntn"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/ntn"
	app.Version = ntn.Version
	app.Usage = "number to number"
	// FIXME: enable autocomplete

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, D",
			EnvVar: "NTN_DEBUG",
			Usage:  "Enable debug mode",
		},
	}
	app.Action = convert
	app.Run(os.Args)
}

func convert(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if len(c.Args()) != 1 {
		return fmt.Errorf("usage: ntn <number>")
	}

	inputStr := c.Args()[0]
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return err
	}

	output := ntn.IntegerToFrench(input)

	fmt.Println(output)
	return nil
}
