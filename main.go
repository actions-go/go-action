package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tjamet/go-github-action-toolkit/core"
)

var now = func() time.Time {
	return time.Now()
}

func runMain() {
	sleep := os.Getenv("INPUT_MILLISECONDS")
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
