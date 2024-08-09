package main

import (
    "fmt"
    "strings"
)

func main(){
    var myStr = []rune("rḗsumḗ") // runes are unicode point number representing characters
//    var myStr = "rḗsumḗ"
    var indexed = myStr[0]
    fmt.Println(myStr)
    fmt.Printf("%v, %T\n", indexed, indexed) // runes are alias for int32
    for i, v := range myStr{
        fmt.Println(i, v)
    }
    fmt.Printf("The length of the string is %v\n", len(myStr))

    var myRune = 'a' // declare rune with single quotes
    fmt.Printf("\nmyRune = %v", myRune)

    var strSlice = []string{"r", "a", "j", "a", "t"}
    var catStr = ""
    for i := range strSlice{
        catStr += strSlice[i] // strings are immutable so with every iteration we are creating a new string which is inefficient...
    }
    fmt.Printf("\n%v\n", catStr)

    var strSlice1 = []string{"r", "a", "j", "a", "t"}
    var strBuilder strings.Builder
    for i := range strSlice1{
        strBuilder.WriteString(strSlice[i])
    }
    var catStr1 = strBuilder.String()
    fmt.Printf("\n%v\n", catStr1)
}
