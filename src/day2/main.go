package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	color_totals := map[string]int{
		"green": 13,
		"blue":  14,
		"red":   12,
	}

	game_number := 1
	sum := 0

	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ": ")
		rounds := strings.Split(games[1], "; ")
		game_possile := true

	RoundLoop:
		for _, round := range rounds {
			sets := strings.Split(round, ", ")
			for _, set := range sets {
				values := strings.Split(set, " ")
				cube_color := values[1]
				cube_count, err := strconv.Atoi(values[0])
				if err != nil {
					panic(err)
				}
				fmt.Printf("Game: %d - %s : %d\n", game_number, cube_color, cube_count)
				if cube_count > color_totals[cube_color] {
					game_possile = false
					break RoundLoop
				}
			}
		}
		if game_possile == true {
			sum += game_number
		}
		game_number += 1
	}

	fmt.Println(sum)

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}
