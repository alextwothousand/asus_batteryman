package commands

import (
	"fmt"

	"github.com/alextwothousand/batteryman/batteryman"
	"github.com/urfave/cli/v2"
)

func Status(c *cli.Context) error {
	status, err := batteryman.GetStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println(status)
	return nil
}
