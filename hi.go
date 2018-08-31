package main

import (
    "fmt"
    "io/ioutil"
     "log"
)

func main() {
	files, err := ioutil.ReadDir("./static")
	
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
    }
}