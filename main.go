package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/actions-go/toolkit/core"
)

var (
	exit = os.Exit
	now  = func() time.Time {
		return time.Now()
	}
)

func main() {
	sleep, ok := core.GetInput("myInput")
	if !ok {
		core.Error("Unable to find required input myInput")
		exit(1)
	}
	core.Debug(fmt.Sprintf("Waiting %s milliseconds", sleep))
	core.Debug(now().String())
	delay, err := strconv.Atoi(sleep)
	if err != nil {
		core.Error(err.Error())
		exit(1)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	core.Debug(now().String())
	core.SetOutput("time", now().String())
	core.Info("All good!")
}
