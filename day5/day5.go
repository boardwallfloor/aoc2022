package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func digitCheck(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	} else {
		return false
	}
}

func parseInstruction(zup string) (int, int, int) {
	tick := 0
	amount := 0
	fro := 0
	to := 0
	for i := 0; i <= len(zup)-1; i++ {
		check := digitCheck(rune(zup[i]))
		if check {
			peek := i
			temp := make([]string, 0)
			for check {
				temp = append(temp, string(zup[peek]))
				if i != len(zup)-1 {
					peek++
					check = digitCheck(rune(zup[peek]))
				} else {
					check = false
				}
			}
			num := strings.Join(temp, "")
			if tick == 0 {
				amount, _ = strconv.Atoi(string(num))
			}
			if tick == 1 {
				fro, _ = strconv.Atoi(string(num))
			}
			if tick == 2 {
				to, _ = strconv.Atoi(string(num))
			}
			tick++
			i = peek
		}
	}
	return amount, fro, to
}

func crane9000(line string, crate [][]rune) {
	zup := line
	zup = zup[5:]
	tick := 0
	amount := 0
	to := 0
	fro := 0

	for i := 0; i <= len(zup)-1; i++ {
		check := digitCheck(rune(zup[i]))
		if check {
			peek := i
			temp := make([]string, 0)
			for check {
				temp = append(temp, string(zup[peek]))
				if i != len(zup)-1 {
					peek++
					check = digitCheck(rune(zup[peek]))
				} else {
					check = false
				}
			}
			num := strings.Join(temp, "")
			if tick == 0 {
				amount, _ = strconv.Atoi(string(num))
			}
			if tick == 1 {
				fro, _ = strconv.Atoi(string(num))
			}
			if tick == 2 {
				to, _ = strconv.Atoi(string(num))
			}
			tick++
			i = peek
		}
	}
	froCrate := crate[fro-1]
	toCrate := crate[to-1]
	for i := 0; i < amount; i++ {
		c := froCrate[len(froCrate)-1]
		toCrate = append(toCrate, c)
		froCrate = append(froCrate[:len(froCrate)-1])
	}
	crate[fro-1] = froCrate
	crate[to-1] = toCrate
}

func main() {
	input, err := os.Open("./day5-test")
	if err != nil {
		fmt.Println(err)
	}
	inputCrate := make([]string, 0)
	crate := make([][]rune, 0)
	isCratePart := true

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && isCratePart {

			isCratePart = false
			index := inputCrate[len(inputCrate)-1]
			count := 0
			for rni, rn := range index {
				if rn != ' ' {
					var temp []rune
					for i := len(inputCrate) - 1; i >= 0; i-- {
						pkg := inputCrate[i]
						if unicode.IsUpper(rune(pkg[rni])) {
							temp = append(temp, rune(pkg[rni]))
						}
					}
					crate = append(crate, temp)
					count++
				}
			}
		}
		if isCratePart {
			inputCrate = append(inputCrate, line)
		} else if line != "" {
			// crane9000(line, crate) // part 1
			zup := line
			zup = zup[5:]

			amount, fro, to := parseInstruction(zup)

			froCrate := crate[fro-1]
			toCrate := crate[to-1]
			pkg := froCrate[len(froCrate)-amount:]
			froCrate = froCrate[:len(froCrate)-amount]

			crate[to-1] = append(toCrate, pkg...)
			crate[fro-1] = froCrate

		}
	}
	for _, line := range crate {
		fmt.Print(string(line[len(line)-1]))
	}
}
