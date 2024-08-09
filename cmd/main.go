package main
import "fmt"

func main(){
    fmt.Println("Hello World!")

    var intNum int
    fmt.Println(intNum)

    var floatNum float64
    fmt.Println(floatNum)

    // cannot perform arithmetic between different types
    // Need to do casting to perform that

    var result float64 = floatNum + float64(intNum)
    fmt.Println(result)

    var intNum1 int = 3
    var intNum2 int = 2
    fmt.Println(intNum1 / intNum2)
    fmt.Println(intNum1 % intNum2)

    var myString string = "Hello" + " " + "World"
    fmt.Println(myString)

    var myBoolean bool = false
    fmt.Println(myBoolean)

    var myStr = "string"
    myStr1 := "string1"
    fmt.Println(myStr)
    fmt.Println(myStr1)

    var1, var2 := 1, 2
    fmt.Println(var1, var2)

    // Cannot change const value
    const myConst string = "const value"
    fmt.Println(myConst)
}
