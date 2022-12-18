package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func copy(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	return err
}

func patchMakefile(fileName, day string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := fmt.Sprintf(`
.PHONY: %s
%s:
	@echo "expected: ?"
	$(run1)
	@echo "expected: ?"
	$(run2)
`, day, day)

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("USAGE: create_day XX"))
	}

	match := regexp.MustCompile(`^day(\d\d)$`).FindStringSubmatch(os.Args[1])
	if match == nil {
		panic(fmt.Errorf("USAGE: create_day XX"))
	}
	day := match[1]

	dayDir := fmt.Sprintf("day%s", day)
	if _, err := os.Stat(dayDir); !os.IsNotExist(err) {
		panic(fmt.Errorf("day directory %s already exists", dayDir))
	}

	if err := os.Mkdir(dayDir, os.ModePerm); err != nil {
		panic(err)
	}

	if err := copy("template/main.go", fmt.Sprintf("%s/main.go", dayDir)); err != nil {
		panic(err)
	}
	if err := copy("template/test.txt", fmt.Sprintf("%s/test.txt", dayDir)); err != nil {
		panic(err)
	}
	if err := copy("template/puzzle.txt", fmt.Sprintf("%s/puzzle.txt", dayDir)); err != nil {
		panic(err)
	}

	patchMakefile("Makefile", day)
}
