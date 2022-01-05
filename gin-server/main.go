package main

import (
	"fmt"
	"os"

	"circlesServer/modules/server"
)

var mode = 0

func main() {
	fmt.Println("Args : ", os.Args)
	if len(os.Args) > 0 && os.Args[1] == `dev` {
		mode = 1
	}
	server.Serve(mode)
}
