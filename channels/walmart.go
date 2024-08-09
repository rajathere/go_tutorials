package main

import (
    "fmt"
    "math/rand"
    "time"
)

var MAX_CHEESE_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main(){
    var cheeseChannel = make(chan string)
    var tofuChannel = make(chan string)
    var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
    for i:= range websites{
        go checkCheesePrices(websites[i], cheeseChannel)
        go checkTofuPrices(websites[i], tofuChannel)
    }
    sendMessage(cheeseChannel, tofuChannel)
}

func checkCheesePrices(website string, cheeseChannel chan string){
    for{
        time.Sleep(time.Second*1)
        var cheesePrice = rand.Float32()*20
        if cheesePrice <= MAX_CHEESE_PRICE{
            cheeseChannel <- website
            break
        }
    }
}

func checkTofuPrices(website string, tofuChannel chan string){
    for{
        time.Sleep(time.Second*1)
        var tofuPrice = rand.Float32()*20
        if tofuPrice <= MAX_TOFU_PRICE{
            tofuChannel <- website
            break
        }
    }
}


func sendMessage(cheeseChannel chan string, tofuChannel chan string){
    select{
        case website := <-cheeseChannel:
            fmt.Printf("Text: Found a deal on cheese at %s\n", website)
       case website := <-tofuChannel:
            fmt.Printf("Email: Found a deal on tofu at %s\n", website)
    }
}
