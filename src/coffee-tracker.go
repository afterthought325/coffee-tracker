package main

import (
    "fmt"
    "menteslibres.net/gosexy/redis"
)

func main() {

    var client  *redis.Client
    var err error

    client = redis.New()

    err = client.Connect("127.0.0.1",6973)

    if err != nil {
        fmt.Println("Connect failed: %s\n")
        return
    }

    fmt.Println("Connected to redis-server.")

  // Get list of Current coffee Items
    //var cl int64 // variable to store length of "coffees"
    cl,err := client.LLen("coffees")
    // grab "coffees" list and store it in "Coffees" array
    Coffees := make([]string,int(cl))
    Coffees,err = client.LRange("coffees",0,-1)
    //print out each of the items
    for n := 0; n <= int(cl); n++ {
        fmt.Println(Coffees[n])
    }

    client.Quit()

}

