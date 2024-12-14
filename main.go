package main

import (
	"fmt"
	"monkeylang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome the the Monkey Console\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
