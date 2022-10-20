package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cwinters8/tools/timer"
)

func cliSetup() error {
	strDuration := flag.String("d", "", "string representation of a duration as defined by the ttimer package (https://github.com/drgrib/ttimer#duration-timing)")
	title := flag.String("t", "", "title of timer")
	flag.Parse()
	if len(*strDuration) < 1 {
		return fmt.Errorf("a string representation of duration at least 1 character long is required")
	}
	duration, err := timer.ParseDuration(*strDuration)
	if err != nil {
		return fmt.Errorf("failed to parse duration string %s: %w", *strDuration, err)
	}
	return timer.Start(duration, *title)
}

func main() {
	if err := cliSetup(); err != nil {
		log.Fatal("failed to setup app: ", err)
	}
}
