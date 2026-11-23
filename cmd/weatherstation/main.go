package main

import (
	"github.com/Golobii/weatherstation/vreme"
	"github.com/urfave/cli/v3"
	"os"
	"log"
	"context"
)

func main() {
	var minTemp float64 
    var maxTemp float64 
    var nSamples int 
    
    cmd := &cli.Command{
        Flags: []cli.Flag{
            &cli.Float64Flag{
                Name:        "minTemp",
                Value:       -50,
                Usage:       "Minimum temperature",
                Destination: &minTemp,
            },
            &cli.Float64Flag{
                Name:        "maxTemp",
                Value:       50,
                Usage:       "Maximum temperature",
                Destination: &maxTemp,
            },
            &cli.IntFlag{
                Name:        "nSamples",
                Value:       3,
                Usage:       "Number of samples needed to calculate statistics",
                Destination: &nSamples,
            },
        },
        
        Action: func(context.Context, *cli.Command) error {
            vreme.Init(minTemp, maxTemp, nSamples)

			// Application
			return nil
        },
    }
    
    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
