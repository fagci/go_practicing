package main

import (
    "flag"
    "fmt"
    "net"
    "strconv"
    "sync"
)

func check(ip string, port string, wg *sync.WaitGroup) {
    defer wg.Done()
    c, e := net.Dial("tcp", ip + ":" + port)
    if e == nil {
        c.Close()
        fmt.Println(port)
    }
}

func main() {
    ip := flag.String("h", "127.0.0.1", "Host to scan")
    flag.Parse()

    var wg sync.WaitGroup
    
    port_from := 1
    port_till := 65536
    wg.Add(port_till - port_from)
    
    for i := port_from; i < port_till; i++ {
        go check(*ip, strconv.Itoa(i), &wg)
    }
    
    wg.Wait()
}
