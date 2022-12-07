define run1
	@go run day$@/main.go part1 < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
	@echo
endef

define run2
	@go run day$@/main.go part2 < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part2 < day$@/puzzle.txt
	@echo
endef

.PHONY: create
create:
	@go run tools/create_day/main.go ${DAY}

.PHONY: all
all: 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24

.PHONY: format
format:
	go fmt helpers/*.go
	go fmt day01/main.go
	go fmt day02/main.go
	go fmt day03/main.go
	go fmt day04/main.go
	go fmt day05/main.go
	go fmt day06/main.go
	go fmt day07/main.go
	go fmt day08/main.go
	go fmt day09/main.go
	go fmt day10/main.go
	go fmt day11/main.go
	go fmt day12/main.go
	go fmt day13/main.go
	go fmt day14/main.go
	go fmt day15/main.go
	go fmt day16/main.go
	go fmt day17/main.go
	go fmt day18/main.go
	go fmt day19/main.go
	go fmt day20/main.go
	go fmt day21/main.go
	go fmt day22/main.go
	go fmt day23/main.go
	go fmt day24/main.go

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

.PHONY: ??
??:
	@echo "expected: ?"
	$(run1)
	@echo "expected: ?"
	$(run2)