package main

import "fmt"
import "math"
import "time"

func main() {
        fmt.Println("Pausing")
        <-time.After(time.Duration(math.MaxInt64))
}
