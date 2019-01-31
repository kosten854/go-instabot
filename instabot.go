package main

import (
	"fmt"
)

func main() {
	// Gets the command line options
	parseOptions()

	// Gets the config
	getConfig(cnfg)

	// Tries to login
	login()

	if *unfollow {
		syncFollowers()
	} else if *run || *interval > 0 {

		// Loop through tags ; follows, likes, and comments, according to the config file
		if *interval == 0 { // Once
			loopTags()
		} else { // From interval

			// Start interval

			stop := make(chan bool)
			go setInterval(loopTags, *interval, stop)

			// Stop control
			for {
				fmt.Println("Write \"stop\" from stop")
				cmd := ""
				fmt.Scan(&cmd)

				if cmd == "stop" {
					// Stop the ticket, ending the interval go routine
					stop <- true
					return
				}

			}
		}
	}
}
