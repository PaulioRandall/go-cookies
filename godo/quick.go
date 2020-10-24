package godo

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	RootDir     string
	BuildDir    string
	BuildPerm   os.FileMode
	ExeFile     string
	BuildFlags  string
	TestTimeout string // E.g. 1m, 5s, 250ms, etc
	MainPkg     string
	Usage       string
}

func (c Config) Godo() (int, error) {

	if len(os.Args) < 2 {
		return EXIT_BAD, fmt.Errorf("Missing command")
	}

	switch cmd := os.Args[1]; cmd {
	case "help":
		fmt.Println(c.Usage)
		return EXIT_OK, nil

	case "clean":
		if _, e := RemoveDir(c.BuildDir); e != nil {
			return EXIT_BAD, e
		}
		return EXIT_BAD, nil

	case "build":
		if e := try(c.Setup, c.Build, c.Fmt); e != nil {
			return EXIT_BAD, e
		}
		return EXIT_OK, nil

	case "test":
		if e := try(c.Setup, c.Build, c.Fmt, c.Test); e != nil {
			return EXIT_BAD, e
		}
		return EXIT_OK, nil

	case "run":
		if e := try(c.Setup, c.Build, c.Fmt, c.Test); e != nil {
			return EXIT_BAD, e
		}
		return c.Run(os.Args[2:]...)

	default:
		return EXIT_BAD, fmt.Errorf("Unknown command: %s", cmd)
	}
}

func (c *Config) Setup() error {
	if _, e := RemoveDir(c.BuildDir); e != nil {
		return e
	}
	if e := os.MkdirAll(c.BuildDir, c.BuildPerm); e != nil {
		return Wrap(e, "Failed to create build directory %s", c.BuildDir)
	}
	return nil
}

func (c Config) Help() {
	fmt.Println(c.Usage)
}

func (c Config) Clean() error {
	_, e := RemoveDir(c.BuildDir)
	return e
}

func (c Config) Build() error {
	g, e := NewGo(c.RootDir)
	if e != nil {
		return e
	}
	return g.Build(c.BuildDir, c.BuildFlags, c.MainPkg)
}

func (c Config) Fmt() error {
	g, e := NewGo(c.RootDir)
	if e != nil {
		return e
	}
	return g.FmtAll()
}

func (c Config) Test() error {
	g, e := NewGo(c.RootDir)
	if e != nil {
		return e
	}
	return g.TestAll(c.TestTimeout)
}

func (c Config) Run(args ...string) (int, error) {
	var e error
	exePath := filepath.Join(c.BuildDir, c.ExeFile)
	if exePath, e = filepath.Abs(exePath); e != nil {
		return EXIT_BAD, Wrap(e, "Couldn't find %s", exePath)
	}
	return Run(exePath, c.BuildDir, args...)
}

func try(funcs ...func() error) error {
	for _, f := range funcs {
		if e := f(); e != nil {
			return e
		}
	}
	return nil
}
