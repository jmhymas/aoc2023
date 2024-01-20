package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(games []string) {
	color_totals := map[string]int{
		"green": 13,
		"blue":  14,
		"red":   12,
	}
	sum := 0

	for game_number, game := range games {
		rounds := strings.Split(game, "; ")
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
			sum += game_number + 1
		}
	}

	fmt.Println(sum)
}

func part2(games []string) {

	power := 0

	for game_number, game := range games {
		rounds := strings.Split(game, "; ")

		max_cube_numbers := map[string]int{
			"green": 0,
			"blue":  0,
			"red":   0,
		}

		for _, round := range rounds {
			sets := strings.Split(round, ", ")
			for _, set := range sets {
				values := strings.Split(set, " ")

				cube_color := values[1]
				cube_count, err := strconv.Atoi(values[0])
				if err != nil {
					panic(err)
				}

				max_cube_number := max_cube_numbers[cube_color]
				if cube_count > max_cube_number {
					max_cube_numbers[cube_color] = cube_count
				}
			}
		}
		product := 1
		for color, number := range max_cube_numbers {
			fmt.Printf("Game %d - %s : %d\n", game_number, color, number)
			product *= number
		}

		power += product
	}

	fmt.Println(power)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var games []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		games = append(games, line[1])

	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	//part1(games)
	part2(games)
}
