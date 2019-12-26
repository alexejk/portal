package portal

import (
	"os"
	"os/exec"
)

func Connect() error {

	// Command:
	// ssh -L <local>:<?host>:<remote> <host>

	c := exec.Command("ssh", "-L", "3000:10.0.3.10:32400", "jaina")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
