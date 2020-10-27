package goquick

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PaulioRandall/go-cookies/gobuild"
)

func Clean(buildDir string) {
	e := os.RemoveAll(buildDir)
	ExitIfErr(e, "Failed to remove build directory: %s", buildDir)
}

func Setup(buildDir string, mode os.FileMode) {
	e := os.MkdirAll(buildDir, mode)
	ExitIfErr(e, "Failed to make build directory: %s", buildDir)
}

func Build(rootDir string, args ...string) {
	g, e := gobuild.NewGo(rootDir)
	ExitIfErr(e, "Failed to build")
	e = g.Build(args...)
	ExitIfErr(e, "Failed to build")
}

func Format(rootDir string, args ...string) {
	g, e := gobuild.NewGo(rootDir)
	ExitIfErr(e, "Failed to format")
	e = g.Fmt(args...)
	ExitIfErr(e, "Failed to format")
}

func Test(rootDir string, args ...string) {
	g, e := gobuild.NewGo(rootDir)
	ExitIfErr(e, "Testing failed")
	e = g.Test(args...)
	ExitIfErr(e, "Testing failed")
}

func Run(buildDir, mainPkgName string, args ...string) int {
	var e error
	exePath := filepath.Join(buildDir, mainPkgName)
	exePath, e = filepath.Abs(exePath)
	ExitIfErr(e, "Failed to run")
	code, e := gobuild.RunCmd(exePath, buildDir, args...)
	ExitIfErr(e, "Failed to run")
	return code
}

func UsageErr(usage, msg string, args ...interface{}) {
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n\n", args...)
	fmt.Println(usage)
	os.Exit(code)
}

func ExitIfErr(cause error, msg string, args ...interface{}) {
	if cause == nil {
		return
	}
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n", args...)
	fmt.Printf("Caused by: %+v", cause)
	os.Exit(code)
}

func AbsPath(rel string) string {
	p, e := filepath.Abs(rel)
	ExitIfErr(e, "Failed to identify path")
	return p
}
