package main

import (
    "fmt"
    "time"
)

func main(){
    // Channels enable goroutines to pass around information
    /* Channels have 3 features
       1. they hold data
       2. they are thread safe
       3. we can listen for data when data is added/removed from channel
    */

//    var c = make(chan int) // this channel can hold a single int value
//    c <- 1 // special way <- to add value to a channel, can be considered as an array holding only 1 value c: [1]
    // c <-1 will create a dealock as we wait here forever for someone to read it, fatal error: all goroutines are asleep - deadlock!
//    var i = <- c // retrieve value from channel, value gets popped from channel (and assigned to i) so channel is empty now
//    fmt.Println(i)


    // Right way to use channels
//    var c = make(chan int)
    var c = make(chan int, 4) // buffered channel
    go process(c) // goroutine prevents us from waiting here
//    for i:=0; i<=5; i++{
    for i:= range c{
        fmt.Println(i) // popping out value of channel
        time.Sleep(time.Second)
    }
}

func process(c chan int){
//    defer close(c) // do this right before the function exists
    for i:=0; i<5; i++{
        c <- i
    }
    close(c) // close the channel after 5 iterations to not wait for read from channel
    fmt.Println("Exiting process")
}
