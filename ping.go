package main

import (
    "fmt"
    "time"
)

func pingSender(pingCh chan <- string) {
    for {
        pingCh <- "ping"
        time.Sleep(time.Second)
    }
}

func main() {
    pingCh := make(chan string)

    go pingSender(pingCh)    
   
    for {
        message := <- pingCh
        fmt.Println(message)
    }
}
