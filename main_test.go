package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {

	if os.Getenv("RUN_MAIN") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMain")
	cmd.Env = append(os.Environ(), "RUN_MAIN=1")
	out, err := cmd.CombinedOutput()
	e, ok := err.(*exec.ExitError)
	require.Equal(t, ok, true)
	require.Equal(t, e.Success(), false)
	require.Contains(t, string(out), "parameters are required")
}
