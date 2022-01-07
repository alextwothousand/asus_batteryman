package main

import (
	"log"
	"os"

	"github.com/alextwothousand/batteryman/cli/commands"
	"github.com/urfave/cli/v2"
)

/*func main() {
	threshold, err := batteryman.GetThreshold()
	if err != nil {
		panic(err)
	}

	fmt.Printf("current battery threshold: %d\n", threshold)

	capacity, err := batteryman.GetCapacity()
	if err != nil {
		panic(err)
	}

	fmt.Printf("current charge levels: %d\n", capacity)

}*/

func main() {
	app := &cli.App{
		Name:  "batteryman",
		Usage: "monitor or configure your computer's battery",
		Commands: []*cli.Command{
			{
				Name:   "status",
				Usage:  "get the current battery charge status",
				Action: commands.Status,
			},
			{
				Name:   "device",
				Usage:  "get the name of the active battery device",
				Action: commands.Device,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
