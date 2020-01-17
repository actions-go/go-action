package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var now = func() time.Time {
	return time.Now()
}

func runMain(w io.Writer) {
	fmt.Println("hello world")
	sleep := os.Getenv("INPUT_MILLISECONDS")
	fmt.Fprintf(w, "::debug::Waiting %s milliseconds\n", sleep)
	fmt.Fprintln(w, "::debug::"+now().String())
	delay, err := strconv.Atoi(sleep)
	if err != nil {
		fmt.Fprintln(w, "::error::"+err.Error())
		return
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Fprintln(w, "::debug::"+now().String())
	fmt.Fprintln(w, "::set-output name=time::"+now().String())
}

func main() {
	runMain(os.Stdout)
}
