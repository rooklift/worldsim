package main

import (
    "fmt"
)

var LogChan chan string

func init() {
    LogChan = make(chan string)
}

func logger() {
    for {
        msg := <- LogChan
        fmt.Printf(msg)
    }
}
