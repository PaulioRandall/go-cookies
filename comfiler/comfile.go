package comfiler

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// Comfile represents a single template along with resources to populate it.
type Comfile struct {
	Template  string // Path to the template file
	Resources string // Path to the root folder of injectable resources
}

// Inject is NOT designed to be called directly, instead it is used in templates
// and called by the "text/template" templating engine.
//
// It takes a 'filename' that is relative to the Resources directory and returns
// its content with each line indented with 'n' tabs. This is used within the
// Template to replace placeholders with some content.
func (com *Comfile) Inject(filename string, n int) string {
	path := com.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s := string(bytes)
	return com.indentEachLine(s, n, "\t")
}

// indentEachLine indents each line of 's' with 'n' instances of 'p'. The
// modified string is then returned. Unix newline '\n' is assumed.
func (com *Comfile) indentEachLine(s string, n int, p string) string {
	prefix := strings.Repeat(p, n)
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = prefix + l
	}
	return strings.Join(lines, "\n")
}

// Compile creates the destination file by copying the template and filling the
// placeholders. Placeholders are relative references to files within the
// the resources directory.
func (com *Comfile) Compile(dst string) error {
	var err error

	t, err := template.ParseFiles(com.Template)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, com)
	if err != nil {
		os.Remove(f.Name())
		return err
	}

	return nil
}
