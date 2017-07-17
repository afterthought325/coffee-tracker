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

    fmt.Println("Welcome to my Coffee Counter. You can see how many times I've made coffee with the given brew type. Then you can increment one when you've made another cup")
  // Get list of Current coffee Items
    cl,err := client.LLen("coffees")// cl is a variable to store length of "coffees"
    // grab "coffees" list and store it in "Coffees" array
    Coffees := make([]string,int(cl))
    Coffees,err = client.LRange("coffees",0,-1)
    //print out each of the items
    var value string
    var n int64
    for n = 0; n < cl; n++ {
        value,err = client.Get(Coffees[n])
        fmt.Println(n+1,") ",Coffees[n]," : ",value," Cups")
    }
    fmt.Println("Do you want to log a cup of coffee?","If so, input the number corresponding to the brewing method then hit enter, otherwise just hit enter")
    var input int
    var input_count int
    input_count,err = fmt.Scanln(&input)
    if input_count != 0 {
        client.Incr(Coffees[input-1])
    }

    client.Quit()

}

