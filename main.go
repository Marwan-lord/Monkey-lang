package main

import (
    "fmt"
    "os"
    "os/user"
    "monkeylang/repl"
)


func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! Welcome the the Monkey Console\n", user.Username)
    repl.Start(os.Stdin, os.Stdout)

}
