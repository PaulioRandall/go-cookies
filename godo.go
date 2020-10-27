package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/go-cookies/goquick"
)

var (
	ROOT          = goquick.AbsPath(".")
	BUILD         = filepath.Join(ROOT, "build")
	BUILD_FLAGS   = ""   // "-gcflags -m -ldflags -s -w"
	TEST_TIMEOUT  = "2s" // E.g. 1m, 5s, 250ms, etc
	MAIN_PKG_NAME = "cmd"
	MAIN_PKG      = "github.com/PaulioRandall/go-cookies/" + MAIN_PKG_NAME
	USAGE         = `Usage:
	help       Show usage
	clean      Remove build files
	build      Build -> format
	test       Build -> format -> test
	run        Build -> format -> test -> run`
)

func main() {

	code := 0
	args := os.Args[1:]

	if len(args) == 0 {
		goquick.UsageErr(USAGE, "Missing command argument")
	}

	switch cmd := args[0]; strings.ToLower(cmd) {
	case "help":
		fmt.Println(USAGE)

	case "clean":
		goquick.Clean(BUILD)

	case "build":
		goquick.Clean(BUILD)
		goquick.Setup(BUILD, os.ModePerm)
		goquick.Build(ROOT, "-o", BUILD, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT, "./...")

	case "test":
		goquick.Clean(BUILD)
		goquick.Setup(BUILD, os.ModePerm)
		goquick.Build(ROOT, "-o", BUILD, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT, "./...")
		goquick.Test(ROOT, "-timeout", TEST_TIMEOUT, "./...")

	case "run":
		goquick.Clean(BUILD)
		goquick.Setup(BUILD, os.ModePerm)
		goquick.Build(ROOT, "-o", BUILD, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT, "./...")
		goquick.Test(ROOT, "-timeout", TEST_TIMEOUT, "./...")
		code = goquick.Run(BUILD, MAIN_PKG_NAME)

	default:
		goquick.UsageErr(USAGE, "Unknown command argument %q", cmd)
	}

	fmt.Printf("\nExit: %d\n", code)
	os.Exit(code)
}
