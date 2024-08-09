package main

import "fmt"


func main(){
    var pointerA *int = new(int) // pointers store addresses of memory locations
    fmt.Printf("Addresses stored by pointer: %v \nValue stored by pointer: %v\n", pointerA, *pointerA)
    *pointerA = 10
    fmt.Printf("Addresses stored by pointer: %v \nValue stored by pointer: %v\n", pointerA, *pointerA)
    var a int
    pointerA = &a
    fmt.Printf("Addresses stored by pointer: %v \nValue stored by pointer: %v\n", pointerA, *pointerA)

    var slice [5]float64 = [5]float64{1,2,3,4,5}
    fmt.Printf("Address of slice in main %p\n", &slice)
    var res = square(&slice)
    fmt.Println(res)

}

func square(thing2 *[5]float64) [5]float64{
    fmt.Printf("Address of slice in square %p\n", thing2) // different address if we pass the array as it copies, same address if we pass address in main and accept pointers here
    for i := range thing2{
        thing2[i] = thing2[i] * thing2[i]
    }
    return *thing2
}
