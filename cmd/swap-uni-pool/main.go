package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Swap via Uniswap pool",
		Usage:  "Swap via SwapRouter with single pool",
		Action: handle,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
}

func handle(*cli.Context) error {
	return nil
}
