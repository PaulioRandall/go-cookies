package quick

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PaulioRandall/go-cookies/go/goexe"
)

// Clean removes the build directory. If an error occurs it is immediately
// printed and the program exits with code 1.
func Clean(buildDir string) {
	e := os.RemoveAll(buildDir)
	ExitIfErr(e, "Failed to remove build directory: %s", buildDir)
}

// Setup creates the build directory and any parents. If an error occurs it is
// immediately printed and the program exits with code 1.
func Setup(buildDir string, mode os.FileMode) {
	e := os.MkdirAll(buildDir, mode)
	ExitIfErr(e, "Failed to make build directory: %s", buildDir)
}

// Build performs 'go build ...' with 'root' as the working directory. If an
// error occurs it is immediately printed and the program exits with code 1.
func Build(root string, args ...string) {
	g, e := goexe.NewGo(root)
	ExitIfErr(e, "Build failed")
	e = g.Build(args...)
	ExitIfErr(e, "Build failed")
}

// Fmt performs 'go fmt ...' with 'root' as the working directory. If an
// error occurs it is immediately printed and the program exits with code 1.
func Fmt(root string, args ...string) {
	g, e := goexe.NewGo(root)
	ExitIfErr(e, "Format failed")
	e = g.Fmt(args...)
	ExitIfErr(e, "Format failed")
}

// Test performs 'go test ...' with 'root' as the working directory. If an
// error occurs it is immediately printed and the program exits with code 1.
func Test(root string, args ...string) {
	g, e := goexe.NewGo(root)
	ExitIfErr(e, "Testing failed")
	e = g.Test(args...)
	ExitIfErr(e, "Testing failed")
}

// Vet performs 'go vet ...' with 'root' as the working directory. If an
// error occurs it is immediately printed and the program exits with code 1.
func Vet(root string, args ...string) {
	g, e := goexe.NewGo(root)
	ExitIfErr(e, "Vet failed")
	e = g.Vet(args...)
	ExitIfErr(e, "Vet failed")
}

// Run executes 'exe' within 'buildDir' returning the exit code. If an
// error occurs it is immediately printed and the program exits with code 1.
func Run(buildDir, exe string, args ...string) int {
	var e error
	exePath := filepath.Join(buildDir, exe)
	exePath, e = filepath.Abs(exePath)
	ExitIfErr(e, "Failed to execute %s", exe)
	code, e := goexe.RunCmd(exePath, buildDir, args...)
	ExitIfErr(e, "Failed to execute %s", exePath)
	return code
}

// Usage error prints the error message, then the program usage, and finally
// exits the program with code 1.
func UsageErr(usage, msg string, args ...interface{}) {
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n\n", args...)
	fmt.Println(usage)
	os.Exit(code)
}

// ExitIfErr prints the error message, then the cause, and finally exits the
// program with code 1 if the cause is not nil else the function returns without
// side effect.
func ExitIfErr(cause error, msg string, args ...interface{}) {
	if cause == nil {
		return
	}
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n", args...)
	fmt.Printf("Caused by: %+v\n", cause)
	os.Exit(code)
}

// AbsPath returns the absolute path of 'rel'.  If an error occurs it is
// immediately printed and the program exits with code 1.
func AbsPath(rel string) string {
	p, e := filepath.Abs(rel)
	ExitIfErr(e, "Failed to identify path")
	return p
}
