package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/go-cookies/go/quick"
)

var (
	ROOT      = quick.AbsPath(".")
	BUILD     = filepath.Join(ROOT, "build")
	PROJ_PATH = "github.com/PaulioRandall/go-cookies"
	MAIN_PKG  = "cmd"
	USAGE     = `Usage:
	help       Show usage
	clean      Remove build files
	build      Build -> format -> vet
	test       Build -> format -> test -> vet
	run        Build -> format -> test -> vet -> run`
)

var (
	BUILD_ARGS = []string{
		"-o", BUILD,
		"", // "-gcflags -m -ldflags -s -w"
		PROJ_PATH + "/" + MAIN_PKG,
	}
	FMT_ARGS  = []string{"./..."}
	TEST_ARGS = []string{"-timeout", "2s", "./..."}
	VET_ARGS  = []string{"./..."}
)

func main() {

	code := 0
	args := os.Args[1:]

	if len(args) == 0 {
		quick.UsageErr(USAGE, "Missing command argument")
	}

	switch cmd := args[0]; strings.ToLower(cmd) {
	case "help":
		fmt.Println(USAGE)

	case "clean":
		quick.Clean(BUILD)

	case "build":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)

	case "test":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Test(ROOT, TEST_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)

	case "run":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Test(ROOT, TEST_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)
		code = quick.Run(BUILD, MAIN_PKG)

	default:
		quick.UsageErr(USAGE, "Unknown command argument %q", cmd)
	}

	fmt.Printf("\nExit: %d\n", code)
	os.Exit(code)
}
