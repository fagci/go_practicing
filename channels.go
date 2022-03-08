package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- fmt.Sprintf("Num: %d", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
        fmt.Println(<- ch)
	}
}
