package main

import (
    "fmt"
    "github.com/infrx0/infrx_test/greetings"
)

func main() {
    greeting := greetings.GetGreeting("World")
    fmt.Println(greeting)
}
