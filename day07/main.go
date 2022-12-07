package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

type Dir struct {
	root         *Dir
	parent       *Dir
	dirs         map[string]*Dir
	files        map[string]int
	computedSize int
}

func (d *Dir) cd(name string) *Dir {
	if name == "/" {
		return d.root
	} else if name == ".." {
		if d.parent != nil {
			return d.parent
		}
		panic(fmt.Errorf("'cd ..': already at /"))
	}

	if subdir, ok := d.dirs[name]; ok {
		return subdir
	}
	panic(fmt.Errorf("'cd %s': subdir does not exist", name))
}

func (d *Dir) addSubdir(name string) {
	if _, found := d.dirs[name]; found {
		panic(fmt.Errorf("'addSubdir %s': subdir already exists", name))
	}
	d.dirs[name] = makeDir(d)
}

func (d *Dir) addFile(name string, size int) {
	if existingSize, found := d.files[name]; found {
		panic(fmt.Errorf("'addFile %s %d': file already exists with size=%d", name, size, existingSize))
	}
	d.files[name] = size
}

func (d *Dir) size() int {
	if d.computedSize < 0 {
		d.computedSize = 0
		for _, subdir := range d.dirs {
			d.computedSize += subdir.size()
		}
		for _, fileSize := range d.files {
			d.computedSize += fileSize
		}
	}
	return d.computedSize
}

func makeDir(parent *Dir) *Dir {
	d := &Dir{nil, parent, make(map[string]*Dir), make(map[string]int), -1}

	if parent == nil {
		d.root = d
	} else {
		d.root = parent.root
	}

	return d
}

func (d *Dir) collectSizes(sizes *[]int) {
	(*sizes) = append(*sizes, d.size())
	for _, subdir := range d.dirs {
		subdir.collectSizes(sizes)
	}
}

func main() {
	reCd := regexp.MustCompile(`^\$ cd (.+)$`)
	reLs := regexp.MustCompile(`^\$ ls$`)
	reDir := regexp.MustCompile(`^dir (.+)$`)
	reFile := regexp.MustCompile(`^(\d+) (.+)$`)

	root := makeDir(nil)
	cwd := root
	helpers.ReadStdin(func(line string) {
		if match := reCd.FindStringSubmatch(line); match != nil {
			cwd = cwd.cd(match[1])
		} else if match := reLs.FindStringSubmatch(line); match != nil {
			// nothing
		} else if match := reDir.FindStringSubmatch(line); match != nil {
			cwd.addSubdir(match[1])
		} else if match := reFile.FindStringSubmatch(line); match != nil {
			cwd.addFile(match[2], helpers.MustParseInt(match[1]))
		} else {
			panic(fmt.Errorf("bad line: %s", line))
		}
	})

	sizes := make([]int, 0)
	root.collectSizes(&sizes)
	if helpers.Part1() {
		sum := 0
		for _, size := range sizes {
			if size <= 100000 {
				sum += size
			}
		}
		fmt.Println(sum)
	} else {
		actualFree := 70000000 - root.size()
		needed := 30000000 - actualFree

		minSize := -1
		for _, size := range sizes {
			if size >= needed && (minSize < 0 || size < minSize) {
				minSize = size
			}
		}
		fmt.Println(minSize)
	}
}
