package run

import (
	"os"
	"os/exec"
)

// Runner is a wrapper around exec.Cmd
type Runner struct {
	c *exec.Cmd
}

func NewRunner(cmd string, args ...string) *Runner {

	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return &Runner{
		c: c,
	}
}

func (r *Runner) Run() error {
	return r.c.Run()
}
