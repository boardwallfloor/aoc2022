package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(file []byte) int {
	var temp string
	tot := 0
	max := 0
	for i := 0; i < len(file); i++ {
		if file[i] != 10 {
			temp += string(file[i])
		} else {
			if i > 0 && file[i-1] == 10 {
				if tot > max {
					max = tot
				}
				tot = 0
			} else {
				c, err := strconv.Atoi(temp)
				if err != nil {
					log.Fatal(err)
				}

				tot += c
				temp = ""

			}
		}
	}
	return max
}

func part2(file []byte) [3]int {
	var temp string
	tot := 0
	var max [3]int
	for i := 0; i < len(file); i++ {
		if file[i] != 10 {
			temp += string(file[i])
		} else {
			if i > 0 && file[i-1] == 10 {
				if tot > max[0] {
					max[2] = max[1]
					max[1] = max[0]
					max[0] = tot
				} else if tot > max[1] {
					max[2] = max[1]
					max[1] = tot
				} else if tot > max[2] {
					max[2] = tot
				}
				tot = 0
			} else {
				c, err := strconv.Atoi(temp)
				if err != nil {
					log.Fatal(err)
				}

				tot += c
				temp = ""

			}
		}
	}
	return max
}

func main() {
	file, err := os.ReadFile("./day1_1")
	if err != nil {
		log.Fatal(err)
	}

	// p1Max := part1(file)
	// fmt.Println(p1Max)
	p2Max := part2(file)
	tot := 0
	for _, v := range p2Max {
		tot += v
	}
	fmt.Println(tot)
}
