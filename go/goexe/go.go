package goexe

import (
	"os"
	"os/exec"

	"github.com/PaulioRandall/go-cookies/cookies"
)

// Go represents a wrapper to the Go compiler. Functionality is provided for
// building, formatting, and testing.
type Go struct {
	Path    string
	WorkDir string
}

// NewGo returns a new Go struct. 'workDir' may be empty to signify the current
// working directory should be used.
func NewGo(workDir string) (Go, error) {

	var e error
	g := Go{WorkDir: workDir}

	if g.WorkDir == "" {
		if g.WorkDir, e = os.Getwd(); e != nil {
			return Go{}, cookies.Wrap(e,
				"Unable to identify current working directory")
		}
	}

	if g.Path, e = exec.LookPath("go"); e != nil {
		return Go{}, cookies.Wrap(e,
			"Can't find compiler. Is it installed? Environment variables set?")
	}
	return g, nil
}

func (g Go) NewCmd(action string, args ...string) *exec.Cmd {
	args = append([]string{action}, args...)
	cmd := exec.Command(g.Path, args...)
	cmd.Dir = g.WorkDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (g Go) Build(args ...string) error {
	cmd := g.NewCmd("build", args...)
	return Run(cmd, "Build failed")
}

func (g Go) Fmt(args ...string) error {
	cmd := g.NewCmd("fmt", args...)
	return Run(cmd, "Format failed")
}

func (g Go) Test(args ...string) error {
	cmd := g.NewCmd("test", args...)
	return Run(cmd, "Testing error")
}

func (g Go) Vet(args ...string) error {
	cmd := g.NewCmd("vet", args...)
	return Run(cmd, "Vet failed")
}

func Run(cmd *exec.Cmd, errMsg string) error {
	if e := cmd.Run(); e != nil {
		return cookies.Wrap(e, "Execution failed")
	}
	return nil
}
