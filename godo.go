package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PaulioRandall/go-cookies/gobuild"
)

// Usage:
//		./godo help
//		./godo clean
//		./godo build
//		./godo test
//		./godo run

var config = gobuild.Config{
	RootDir:     absPath("."),
	BuildDir:    filepath.Join(absPath("."), "build"),
	BuildPerm:   os.ModePerm,
	ExeFile:     "cmd",
	BuildFlags:  "",   // "-gcflags -m -ldflags -s -w"
	TestTimeout: "2s", // E.g. 1m, 5s, 250ms, etc
	MainPkg:     "github.com/PaulioRandall/go-cookies/cmd",
	Usage: `Usage:
	help       Show usage
	clean      Remove build files
	build      Build -> format
	test       Build -> format -> test
	run        Build -> format -> test -> run`,
}

func main() {
	exitCode, e := config.Godo()
	if e != nil {
		fmt.Printf("%v\n", e)
	}
	fmt.Printf("Exitcode: %d\n", exitCode)
	os.Exit(exitCode)
}

func absPath(rel string) string {
	p, e := filepath.Abs(rel)
	if e != nil {
		panic(e)
	}
	return p
}
