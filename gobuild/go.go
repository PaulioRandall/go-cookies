package gobuild

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

func (g Go) NewCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(g.Path, args...)
	cmd.Dir = g.WorkDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (g Go) Build(dir, flags, pkg string) error {
	var cmd *exec.Cmd
	if dir == "" {
		cmd = g.NewCmd("build", flags, pkg)
	} else {
		cmd = g.NewCmd("build", "-o", dir, flags, pkg)
	}
	return g.run(cmd, "Build failed")
}

func (g Go) Fmt(pkg string) error {
	cmd := g.NewCmd("fmt", pkg)
	return g.run(cmd, "Format failed")
}

func (g Go) FmtAll() error {
	return g.Fmt("./...")
}

func (g Go) Test(pkg string, timeout string) error {
	var cmd *exec.Cmd
	if timeout == "" {
		cmd = g.NewCmd("test", pkg)
	} else {
		cmd = g.NewCmd("test", pkg, "-timeout", timeout)
	}
	return g.run(cmd, "Testing error")
}

func (g Go) TestAll(timeout string) error {
	return g.Test("./...", timeout)
}

func (g Go) run(cmd *exec.Cmd, errMsg string) error {
	if e := cmd.Run(); e != nil {
		return cookies.Wrap(e, "Execution failed")
	}
	return nil
}
