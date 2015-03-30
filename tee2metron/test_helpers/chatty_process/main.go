package main

import (
	"fmt"
	"os"
	"time"
	"io"
)

func main() {
	done := make(chan struct{})

	go pollStdin(os.Stdin, done)

	loop:
	for {
		fmt.Println("loop")
		fmt.Fprintf(os.Stdout, "Hi from stdout. My args are: %s\n", os.Args[1:])
		fmt.Fprintln(os.Stderr, "Oopsie from stderr")
		time.Sleep(100 * time.Millisecond)

		select {
		case <- done:
			close(done)
			break loop
		default:
		}
	}

	panic("wtf")

}

func pollStdin(stdinReader io.ReadCloser, done chan struct{}) () {
	defer stdinReader.Close()

	fmt.Fscanln(os.Stdin)
	done <- struct{}{}
}