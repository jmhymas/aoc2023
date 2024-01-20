# Advent of Code 2023

## Summary

Advent of Code 2023 solutions in Golang

## Layout

All solutions are located in src/ organized by day

```
./src/day1/main.go

./src/day2/main.go

...

```

## Building/Running Solutions

To run the solutions I usually place a build/ directory in each day and run the following command from within a given days directory

```bash
go build -o ./build/ && ./build/day2
```

Replace day2 with the appropriate value. The .gitignore file has a line so the build directories and their binaries don't get added to the repo.

## Notes

You'll probably see some repeated code in some of the solutions. I want to keep both part1 and part2 in one file for each day, and thanks to the way Advent of Code is structured, sometimes it's easier to just copy and modify the code from part1 to do part2. The main focus here is to have a little fun learning Go and solve the problems at hand, not to spend a bunch of time formatting/refactoring the code.