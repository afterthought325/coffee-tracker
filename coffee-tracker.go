package main

import (
    "fmt"
    "menteslibres.net/gosexy/redis"
)

func main() {

    var client  *redis.Client
    var err error

    client = redis.New()

    err = client.Connect("127.0.0.1",6379)

    if err != nil {
        fmt.Println("Connect failed: %s\n")
        return
    }

    fmt.Println("Connected to redis-server.")

  // Get list of Current coffee Items
    cl,err := client.LLen("coffees")// cl is a variable to store length of "coffees"
    // grab "coffees" list and store it in "Coffees" array
    Coffees := make([]string,int(cl))
    Coffees,err = client.LRange("coffees",0,-1)
    //print out each of the items
    var n int64 
    for n = 0; n <= cl; n++ {
        fmt.Println(Coffees[n])
    }

    client.Quit()

}

