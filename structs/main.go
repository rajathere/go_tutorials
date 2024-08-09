package main

import "fmt"

type gasEngine struct{
    mpg int32
    gallons int32
}

type electricEngine struct{
    mpkwh int32
    kwh int32
}

type owner struct{
    name string
}

type gasEngineOwner struct{
    engine gasEngine
    ownerInfo owner
}

func (e gasEngine) milesLeft() int32 { // assign this method to gasEngine type
    return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() int32 { // assign this method to electricEngine type
    return e.mpkwh * e.kwh
}

type engine interface{
    milesLeft() int32 // the requirement of interface is the struct should have milesLeft method
}

func canMakeIt(e engine, miles int32){
    if miles <= e.milesLeft(){
        fmt.Printf("Can make it!\n")
    }else{
        fmt.Printf("Need to fuel up first\n")
    }
}

func main(){
    var myEngine gasEngine = gasEngine{mpg: 25, gallons: 30}
    fmt.Println(myEngine.mpg, myEngine.gallons)
    var myEngine1 gasEngine = gasEngine{20, 25} // values are assigned in order
    fmt.Println(myEngine1.mpg, myEngine1.gallons)

    var myEngine2 gasEngineOwner = gasEngineOwner{gasEngine{25, 30}, owner{"Rajat"}}
    fmt.Println(myEngine2.engine.mpg, myEngine2.engine.gallons, myEngine2.ownerInfo.name)
    
    fmt.Println(myEngine.mpg * myEngine.gallons)
    fmt.Printf("Total miles left: %v\n", myEngine.milesLeft())

    canMakeIt(myEngine, 52)

    var eEngine electricEngine = electricEngine{30, 40}

    fmt.Printf("Total miles left: %v\n", eEngine.milesLeft())
    canMakeIt(eEngine, 100)

}
