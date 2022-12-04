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

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("USAGE: create_day XX"))
	}

	day := os.Args[1]
	if !regexp.MustCompile(`^\d\d$`).MatchString(day) {
		panic(fmt.Errorf("USAGE: create_day XX"))
	}

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
}
