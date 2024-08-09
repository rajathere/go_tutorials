package main

import(
    "fmt"
    "time"
    "sync"
)

//var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData []int = []int{1,2,3,4,5}
var results []int

func main(){
    // IO Intensive
    t0 := time.Now()
    for i,_ := range dbData{
        wg.Add(1)
        go dbCall(i) // goroutines free up cpu when waiting for another task
    }
    wg.Wait()
    fmt.Printf("Total execution time: %v\n", time.Since(t0))
    fmt.Printf("Results are: %v\n", results)

    // CPU intensive
    t1 := time.Now()
    for i:=0; i<100; i++{
        wg.Add(1)
        go count() // goroutines free up cpu when waiting for another task
    }
    wg.Wait()
    fmt.Printf("Total execution time: %v\n", time.Since(t1))

}

func dbCall(i int){
    var delay float32 = 2000
    time.Sleep(time.Duration(delay)*time.Millisecond)
//    fmt.Printf("Result from db is %v\n", dbData[i])
    save(dbData[i])
    log()
    wg.Done()
}

func save(result int){
    m.Lock()
    results = append(results, result)
    m.Unlock()
}

func log(){
    m.RLock()
    fmt.Printf("The current results are: %v\n", results)
    m.RUnlock()
}

func count(){
    var res int
    for i:=0; i<1000000; i++{
        res += 1
    }
    wg.Done()
}
