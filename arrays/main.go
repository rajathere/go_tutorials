package main
import "fmt"

func main(){
    var intArr [3]int32
    intArr [1] = 132
    fmt.Println(intArr[0])
    fmt.Println(intArr[1:3])

    fmt.Println(&intArr[0])
    fmt.Println(&intArr[1])
    fmt.Println(&intArr[2])

    var floatArr [3]float32 = [3]float32{1.1,2.2,3.3}
    fmt.Println(floatArr)

    vals := [...]int32{1,2,3}
    fmt.Println(vals)

    // slices
    var intSlice []int32 = []int32{4,5,6}
    fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // 3, 3
    intSlice = append(intSlice, 7)
    fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // 4, 6
    // fmt.Printf(intSlice[4]) -- error

    var intSlice2 []int32 = []int32{8,9,10}
    intSlice = append(intSlice, intSlice2...)
    fmt.Println(intSlice)
    fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // 7, 12

    var intSlice3 []int32 = make([]int32, 3, 8) // type, length, capacity
    fmt.Println(intSlice3)

    var myMap map[string]int32 = make(map[string]int32)
    myMap["ID"] = 1
    fmt.Println(myMap)
    var myMap2 = map[string]int32{"Rajat": 29, "Vipin": 31}
    fmt.Println(myMap2)

    var age1, exists1 = myMap2["Rajat"]
    var age2, exists2 = myMap2["Shekhar"]

    fmt.Println(age1, exists1)
    fmt.Println(age2, exists2)

    delete(myMap2, "Rajat")

    var age3, exists3 = myMap2["Rajat"]
    fmt.Println(age3, exists3)

    // loops
    for name, age := range myMap2{
        fmt.Println(name, age)
    }

    for i, v :=range intSlice{
        fmt.Println(i, v)
    }

    // while using for loop
    var i int = 0
    for {
        if i >= 10{
            break
        }
        fmt.Println(i)
        i = i + 1
    }

    for i:=0; i<10; i++{
        fmt.Println(i)
    }

}
