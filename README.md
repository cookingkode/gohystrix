# gohystrix
## Simple wrapper for Hystrix Command pattern

This project is inspired by the Netflix Hystrix project :http://techblog.netflix.com/2012/02/fault-tolerance-in-high-volume.html

The initial commit contains implementation for the Command pattern. Once can define a function, timeout and a fallback. 

***Usage: **
```
package main

import (
    "fmt"
    "gohystrix"
    "time"
)

func mainFunction(param interface{}) interface{} {
    d := param.(int)

    time.Sleep(time.Second * 3) // Vary this to force/not timeout
    fmt.Println("Main Function ", d)

    return d * 2
}

func fallback(param interface{}) interface{} {
    d := param.(int)
    fmt.Println("FALLBACK ", d)

    return d * 4
}

func main() {

    cmd := gohystrix.NewCommand(mainFunction, fallback, 2000)
    res := cmd.Run(2)
    fmt.Println(res)

}
```
