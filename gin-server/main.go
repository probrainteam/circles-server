package main

import (
	"fmt"
	"os"

	"circlesServer/modules/server"
)

var mode = 0

func main() {
	fmt.Println("Args : ", os.Args)
	opt := "deploy"
	if len(os.Args) == 2 {
		opt = os.Args[1]
	}

	fmt.Println("opt : ", opt)
	server.Serve(opt)
}
