package main

import (
    "fmt"
    "os"
    "os/user"
    "go-interpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! Type in commands to start\n", user.Username)
    repl.Start(os.Stdin, os.Stdout)
}
