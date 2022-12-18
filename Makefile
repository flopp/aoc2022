define run1
	@go run day$@/main.go part1 test < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part1 puzzle < day$@/puzzle.txt
	@echo
endef

define run2
	@go run day$@/main.go part2 test < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part2 puzzle < day$@/puzzle.txt
	@echo
endef

all:
	@echo "Run 'make dayXX' to create a template directory for day XX"
	@echo "Run 'make XX' to run the test and puzzle inputs on the day XX solution"

day%:
	@go run tools/create_day/main.go $@

.PHONY: 01
01:
	@echo "expected: 24000"
	$(run1)
	@echo "expected: 45000"
	$(run2)

.PHONY: 02
02:
	@echo "expected: 15"
	$(run1)
	@echo "expected: 12"
	$(run2)

.PHONY: 03
03:
	@echo "expected: 157"
	$(run1)
	@echo "expected: 70"
	$(run2)

.PHONY: 04
04:
	@echo "expected: 2"
	$(run1)
	@echo "expected: 4"
	$(run2)

.PHONY: 05
05:
	@echo "expected: CMZ"
	$(run1)
	@echo "expected: MCD"
	$(run2)

.PHONY: 06
06:
	@echo "expected: 7, 5, 6, 10, 11"
	$(run1)
	@echo "expected: 19, 23, 23, 29, 26"
	$(run2)

.PHONY: 07
07:
	@echo "expected: 95437"
	$(run1)
	@echo "expected: 24933642"
	$(run2)

.PHONY: 08
08:
	@echo "expected: 21"
	$(run1)
	@echo "expected: 16 (not directly given in puzzle description)"
	$(run2)

.PHONY: 09
09:
	@echo "expected: 13"
	$(run1)
	@echo "expected: 1"
	@go run day09/main.go part2 < day09/test.txt
	@echo "expected: 36"
	@go run day09/main.go part2 < day09/test2.txt
	@echo "=>"
	@go run day09/main.go part2 < day09/puzzle.txt

.PHONY: 10
10:
	@echo "expected: 13140"
	$(run1)
	@echo "expected: \n##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n---"
	$(run2)

.PHONY: 11
11:
	@echo "expected: 10605"
	$(run1)
	@echo "expected: 2713310158"
	$(run2)

.PHONY: 12
12:
	@echo "expected: 31"
	$(run1)
	@echo "expected: 29"
	$(run2)

.PHONY: 13
13:
	@echo "expected: 13"
	$(run1)
	@echo "expected: 140"
	$(run2)

.PHONY: 14
14:
	@echo "expected: 24"
	$(run1)
	@echo "expected: 93"
	$(run2)

.PHONY: 15
15:
	@echo "expected: 26"
	$(run1)
	@echo "expected: 56000011"
	$(run2)

.PHONY: 16
16:
	@echo "expected: 1651"
	$(run1)
	@echo "expected: 1707"
	$(run2)

.PHONY: 17
17:
	@echo "expected: 3068"
	$(run1)
	@echo "expected: ?"
	$(run2)

.PHONY: 18
18:
	@echo "expected: 64"
	$(run1)
	@echo "expected: 58"
	$(run2)
