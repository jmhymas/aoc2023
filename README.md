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

You'll probably see some repeated code in some of the solutions. Sometimes it's easier to just copy part 1 and modify it, and I'm more interested in solving the problems then making the code 100% perfect.