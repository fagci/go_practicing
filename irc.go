package main

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
    "time"
)

func read_messages(c net.Conn) {
	for {
		m, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println(m)
	}
}

func auth(c net.Conn) {
    c.Write([]byte("USER g0-b07 * g0-b07 g0-b07\r\n\r\n"))
    c.Write([]byte("NICK g0-b07\r\n\r\n"))
    c.Write([]byte("JOIN #bash\r\n\r\n"))
}

func main() {
	addr := "irc.libera.chat:6667"
	c, _ := net.Dial("tcp", addr)

	go read_messages(c)

    time.Sleep(time.Second*3)

    go auth(c)


	runtime.Goexit()
}
