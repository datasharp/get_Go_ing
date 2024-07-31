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
    message, err := greetings.Hello("Gladys")
    if err != nil {
        log.Fatal(err) // fatal: print error and stop function
    }
    // no error, then print
    fmt.Println(message)

    message, err = greetings.Goodbye("")
    if err != nil {
        log.Fatal(err) 
    }
    fmt.Println(message)
}
