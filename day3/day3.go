package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(file []byte) int {
	rsByte := file
	var temp []byte
	tot := 0
	for i, v := range rsByte {
		if v == 10 || len(rsByte)-1 == i {
			rs1 := temp[0 : len(temp)/2]
			rs2 := temp[len(temp)/2:]

			var d string
			for _, b := range rs1 {
				val := strings.ContainsRune(string(rs2), rune(b))
				if val && !strings.Contains(d, string(b)) {
					// fmt.Println(string(b), b)
					d += string(b)
					// lowercase: 96, uppercase : 38
					if b >= byte(96) {
						tot += int(b) - 96
						// fmt.Println(int(b) - 96)
					}
					if b >= byte(38) && b < byte(96) {
						tot += int(b) - 38
						// fmt.Println(int(b) - 38)
					}
				}
			}

			// for _, b := range rs1 {
			// 	m[b] = true
			// }
			// for _, b := range rs2 {
			// 	_, ex := m[b]
			// 	if ex {
			// 		fmt.Println(string(b))
			// 	}
			// }
			temp = temp[:0]
		} else {
			temp = append(temp, v)
		}
	}
	return tot
}

func part2(file []byte) int {
	rsByte := file
	var temp []byte
	tot := 0
	count := 0
	prev := ""

	for i, v := range rsByte {
		if v == 10 || len(rsByte)-1 == i {
			// count 0 : save current temp string
			if count == 0 {
				prev = string(temp)
			}
			// count 1 : get duplicate from prev saved temp string, save duplicate to another
			if count == 1 {
				d := ""
				for _, rn := range temp {
					if strings.ContainsRune(prev, rune(rn)) && !strings.ContainsRune(d, rune(rn)) {
						d += string(rn)
					}
				}
				prev = d
			}
			// count 2 : get duplicate from current temp compared to duplicate from count 1
			if count == 2 {
				count = 0
				d := ""
				for _, rn := range temp {
					if strings.ContainsRune(prev, rune(rn)) && !strings.ContainsRune(d, rune(rn)) {
						d += string(rn)
						if rn >= byte(96) {
							tot += int(rn) - 96
						}
						if rn >= byte(38) && rn < byte(96) {
							tot += int(rn) - 38
						}
					}
				}
			} else {
				count++
			}
			temp = temp[:0]
		} else {
			temp = append(temp, v)
		}
	}
	return tot
}

func main() {
	file, err := os.ReadFile("./day3")
	if err != nil {
		fmt.Println(err)
	}

	tot := part2(file)
	fmt.Println(tot)
}
