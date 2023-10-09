package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/actions-go/toolkit/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	errorOutput = `::debug::Did not find the input using plain gha input format%2C trying the actions-go one
::error::Unable to find required input myInput
`
	successOutput = `::debug::Waiting 10 milliseconds
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
`
)

func TestRunMain(t *testing.T) {
	t.Run("When the input is not available", func(t *testing.T) {
		exitMagic := "test-run-main"
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, exitMagic, r)
			}
		}()
		t.Cleanup(func() {
			exit = os.Exit
		})
		exit = func(code int) {
			assert.Equal(t, 1, code)
			panic(exitMagic)
		}

		w := bytes.NewBuffer(nil)
		core.SetStdout(w)
		main()

		assert.Equal(t, errorOutput, strings.Replace(w.String(), "%3A", ":", -1))

	})

	t.Run("When the input is available", func(t *testing.T) {
		fd, err := os.CreateTemp("", "githuboutput")
		require.NoError(t, err)

		outputPath := fd.Name()
		os.Setenv("GITHUB_OUTPUT", outputPath)
		t.Cleanup(func() {
			fd.Close()
			os.Remove(outputPath)

		})
		os.Setenv("INPUT_MYINPUT", "10")
		exit = func(code int) {
			assert.Equal(t, 0, code)
		}
		now = func() time.Time {
			return time.Date(2020, 01, 10, 20, 10, 20, 1, time.UTC)
		}

		w := bytes.NewBuffer(nil)
		core.SetStdout(w)
		main()

		assert.Equal(t, successOutput, strings.Replace(w.String(), "%3A", ":", -1))
	})
}
