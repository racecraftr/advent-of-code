// Package skeleton makes skeletons to be filled out with solutions.
package skeleton

import (
	"adventOfCode/util"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.go
var fs embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fs, "templates/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	mainFilename := filepath.Join(util.Dirname(), "../", fmt.Sprintf("y%d/day%02d/main.go", year, day))
	inFilename := filepath.Join(util.Dirname(), "../", fmt.Sprintf("y%d/day%02d/in.txt", year, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	_, err = os.Create(inFilename)
	if err != nil {
		log.Fatalf("creating in.txt file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	fmt.Printf("templates made for %d-day%d\n", year, day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
