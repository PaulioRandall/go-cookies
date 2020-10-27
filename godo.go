package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/go-cookies/goquick"
)

var (
	ROOT_DIR      = goquick.AbsPath(".")
	BUILD_DIR     = filepath.Join(ROOT_DIR, "build")
	BUILD_MODE    = os.ModePerm
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
		goquick.Clean(BUILD_DIR)

	case "build":
		goquick.Clean(BUILD_DIR)
		goquick.Setup(BUILD_DIR, BUILD_MODE)
		goquick.Build(ROOT_DIR, BUILD_DIR, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT_DIR)

	case "test":
		goquick.Clean(BUILD_DIR)
		goquick.Setup(BUILD_DIR, BUILD_MODE)
		goquick.Build(ROOT_DIR, BUILD_DIR, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT_DIR)
		goquick.Test(ROOT_DIR, TEST_TIMEOUT)

	case "run":
		goquick.Clean(BUILD_DIR)
		goquick.Setup(BUILD_DIR, BUILD_MODE)
		goquick.Build(ROOT_DIR, BUILD_DIR, BUILD_FLAGS, MAIN_PKG)
		goquick.Format(ROOT_DIR)
		goquick.Test(ROOT_DIR, TEST_TIMEOUT)
		code = goquick.Run(BUILD_DIR, MAIN_PKG_NAME)

	default:
		goquick.UsageErr(USAGE, "Unknown command argument %q", cmd)
	}

	fmt.Printf("\nExit: %d\n", code)
	os.Exit(code)
}
