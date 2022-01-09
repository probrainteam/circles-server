package main

import (
	"fmt"
	"os"

	"circlesServer/modules/server"
)

var mode = 0

func main() {
	fmt.Println("Args : ", os.Args)
	if len(os.Args) > 0 {
		opt := os.Args[1]
		if opt == `dev` { // use mock data
			mode = 1
		} else if opt == `debug` { // log everything
			mode = 2
		} else {
			panic(fmt.Errorf(`Unknown command : ` + os.Args[1])) // exception
		}
	}
	server.Serve(mode)
}
