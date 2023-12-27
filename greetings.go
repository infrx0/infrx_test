package main

import (
    "fmt"
    "github.com/infrx0/infrx_test/constants"
)

func GetGreeting(name string) string {
    return fmt.Sprintf("%s%s!", constants.GreetingPrefix, name)
}
