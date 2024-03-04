package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		for i := range 4 {
			msg := fmt.Sprintf("message #%d", i)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	for {
		msg, ok := <-ch
		if !ok {
			break
		}

		fmt.Println(msg)
	}
}
