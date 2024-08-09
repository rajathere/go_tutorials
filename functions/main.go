package main
import (
    "errors"
    "fmt"
)

func main(){
    var printValue string = "Hello"
    printMe(printValue)
    var numerator int = 6
    var denominator int = 56 
    var res, remainder, err = intDivision(numerator, denominator)
    switch{
        case err!=nil:
            fmt.Printf(err.Error() + "\n")
        case remainder==0:
            fmt.Printf("The result of the integer division is %v\n", res)
        default:
            fmt.Printf("The result of integer divison is %v and remainder is %v\n", res, remainder)
    }

    switch remainder{
        case 0:
            fmt.Printf("The division was exact\n")
        case 1,2:
            fmt.Printf("The division was close\n")
        default:
            fmt.Printf("The division was not close\n")
    }
/*    if err != nil{
        fmt.Printf(err.Error())
    }else if remainder == 0{
        fmt.Printf("The result of the integer division is %v", res)
    }else{
    fmt.Printf("The result of integer divison is %v and remainder is %v", res, remainder)
    }
*/
}

func printMe(printValue string){
    fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error){
    var err error
    if denominator==0{
        err = errors.New("Cannot Divide By Zero")
        return 0, 0, err
    }
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder, err
}
