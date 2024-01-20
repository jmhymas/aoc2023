package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(games []string) int {
	fmt.Println("PART1")
	fmt.Println("-------------------------------------")
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
				if cube_count > color_totals[cube_color] {
					fmt.Printf("Game %d not possible with %s cube count of %d.\n", game_number+1, cube_color, cube_count)
					game_possile = false
					break RoundLoop
				}
			}
		}
		if game_possile == true {
			fmt.Printf("Game %d possible.\n", game_number+1)
			sum += game_number + 1
		}
	}
	fmt.Println("-------------------------------------")
	return sum
}

func part2(games []string) int {
	fmt.Println("PART2")
	fmt.Println("-------------------------------------")
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

		fmt.Printf("Game %d max values: (", game_number+1)
		for color, number := range max_cube_numbers {
			fmt.Printf("%s : %d, ", color, number)
			product *= number
		}
		fmt.Printf(")\n")
		power += product
	}
	fmt.Println("-------------------------------------")
	return power
}

func main() {
	file, err := os.Open("./input_example.txt")
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

	sum := part1(games)
	power := part2(games)

	fmt.Printf("The sum of the possible games in part 1 is %d\n", sum)
	fmt.Printf("the power of all the required cube numbers in part 2 is %d\n", power)
}
