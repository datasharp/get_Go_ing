package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // set logger properties
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // request greeting message
    message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err) // fatal: print error and stop function
    }
    // no error, then print
    fmt.Println(message)

    message = greetings.Goodbye("Gladys")
    fmt.Println(message)
}
