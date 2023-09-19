package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/actions-go/toolkit/core"
)

var now = func() time.Time {
	return time.Now()
}

func runMain() {
	sleep, ok := core.GetInput("myInput")
	if !ok {
		core.Error("Unable to find required input myInput")
	}
	fmt.Println(sleep)
	core.Debug(fmt.Sprintf("Waiting %s milliseconds", sleep))
	core.Debug(now().String())
	delay, err := strconv.Atoi(sleep)
	if err != nil {
		core.Error(err.Error())
		return
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	core.Debug(now().String())
	core.SetOutput("time", now().String())
}

func main() {
	runMain()
}
