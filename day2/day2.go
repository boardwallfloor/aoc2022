package main

import (
	"fmt"
	"os"
)

func part1(stratByte []byte) int {
	matchup := [][]int{
		// rock, paper, scissor
		{4, 8, 3}, // rock
		{1, 5, 9}, // paper
		{7, 2, 6}, // scissor
	}
	se := 0
	op := 0
	tot := 0
	for i, v := range stratByte {
		switch v {
		case 65:
			op = 0
		case 66:
			op = 1
		case 67:
			op = 2
		case 88:
			se = 0
		case 89:
			se = 1
		case 90:
			se = 2
		}
		if v == 10 || len(stratByte)-1 == i {
			tot += matchup[op][se]
		}
	}
	return tot
}

func part2(stratByte []byte) int {
	se := 0
	op := 0
	tot := 0
	for i, v := range stratByte {
		switch v {
		case 65:
			op = 0
		case 66:
			op = 1
		case 67:
			op = 2
		case 88:
			se = 0
		case 89:
			se = 1
		case 90:
			se = 2
		}
		if v == 10 || len(stratByte)-1 == i {
			if se == 0 {
				switch op {
				case 0:
					tot += 3
				case 1:
					tot += 1
				case 2:
					tot += 2
				}
			}
			if se == 1 {
				tot += op + 1 + 3
			}
			if se == 2 {
				switch op {
				case 0:
					tot += 2
				case 1:
					tot += 3
				case 2:
					tot += 1
				}
				tot += 6
			}

		}
	}
	return tot
}

func main() {
	stratByte, err := os.ReadFile("./day2")
	if err != nil {
		fmt.Println(err)
	}

	p1tot := part1(stratByte)
	fmt.Println(p1tot)
	p2tot := part2(stratByte)
	fmt.Println(p2tot)
}
