package commands

import (
	"fmt"

	"github.com/alextwothousand/batteryman/batteryman"
	"github.com/urfave/cli/v2"
)

func Device(c *cli.Context) error {
	device, err := batteryman.GetBatteryDevice()
	if err != nil {
		panic(err)
	}

	fmt.Println(device)
	return nil
}
