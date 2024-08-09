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

type car[T gasEngine | electricEngine] struct{
    carMake string
    carModel string
    engine T
}

func main(){
    // Generics for int
    var intSlice = []int{1,2,3}
    fmt.Println(sumSlice[int](intSlice))

    // Generics for float32
    var float32Slice = []float32{1.1, 2.2, 3.3}
    fmt.Println(sumSlice[float32](float32Slice))

    // Generics for float64
    var float64Slice = []float64{1.11, 2.22, 3.33}
    fmt.Println(sumSlice[float64](float64Slice))

    // Generics for any type
    var emptySlice = []int{}
    fmt.Println(isEmpty[int](emptySlice))

    var nonEmptySlice = []float32{1.1}
    fmt.Println(isEmpty[float32](nonEmptySlice))

    // Generics for struct
    var gasCar = car[gasEngine]{"Honda", "Civic", gasEngine{15, 30}}
    var electricCar = car[electricEngine]{"Tesla", "Model 3", electricEngine{20, 40}}

    fmt.Println(gasCar)
    fmt.Println(electricCar)

}

func sumSlice[T int | float32 | float64](slice []T) T{
    var sum T
    for _, v := range slice{
        sum += v
    }
    return sum
}

func isEmpty[T any](slice []T) bool{
    return len(slice)==0
}


