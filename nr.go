package main

import (
    "encoding/binary"
    "fmt"
    "math/rand"
    "net"
    "strconv"
    "runtime"
    "time"
)

func random_ip() string {
    for {
        intip := uint32(rand.Float32()*0xD0000000) + 0xFFFFFF
        if !(
            (intip >= 0x0A000000 && intip <= 0x0AFFFFFF) ||
            (intip >= 0x64400000 && intip <= 0x647FFFFF) ||
            (intip >= 0x7F000000 && intip <= 0x7FFFFFFF) ||
            (intip >= 0xA9FE0000 && intip <= 0xA9FEFFFF) ||
            (intip >= 0xAC100000 && intip <= 0xAC1FFFFF) ||
            (intip >= 0xC0000000 && intip <= 0xC0000007) ||
            (intip >= 0xC00000AA && intip <= 0xC00000AB) ||
            (intip >= 0xC0000200 && intip <= 0xC00002FF) ||
            (intip >= 0xC0A80000 && intip <= 0xC0A8FFFF) ||
            (intip >= 0xC6120000 && intip <= 0xC613FFFF) ||
            (intip >= 0xC6336400 && intip <= 0xC63364FF) ||
            (intip >= 0xCB007100 && intip <= 0xCB0071FF) ||
            (intip >= 0xF0000000 && intip <= 0xFFFFFFFF)){
            ip := make(net.IP, 4)
            binary.BigEndian.PutUint32(ip, intip)
            return ip.String()

        }
    }
}

func check(ip string, port string) {
    d := net.Dialer{Timeout: time.Second}
    addr := ip + ":" + port
    c, e := d.Dial("tcp", addr)
    if e == nil {
        c.Close()
        fmt.Println(addr)
    }
}

func worker() {
    for {
        ip := random_ip()
        go check(ip, strconv.Itoa(80))
    }
}

func main() {
    for i := 0; i < 1024; i++ {
        go worker()
    }

    runtime.Goexit()

}
