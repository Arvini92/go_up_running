package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"encoding/json"
    //"strings"
    //"math/big"
)

func main() {

    url := "http://service.explorecalifornia.org/json/tours.php"
    content := contentFromServer(url)
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Response type: %T\n", resp)

    defer resp.Body.close()

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    content := string(bytes)
    fmt.Print(content)

}