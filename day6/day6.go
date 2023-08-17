package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 	input := `bvwbjplbgvbhsrlpgdmjqwftvncz
	// nppdvjthqldpwncqszvftbrmjlhg
	// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
	// zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
	// lineOpen strings.Split(input, "\n")
	// length := 4 // part 1
	length := 14 // part 2
	input, err := os.ReadFile("./day6")
	if err != nil {
		log.Fatal(err)
	}
	line := string(input)
	// for _, line := range lines {
	fmt.Println(line)
	window := make(map[string]int, length)
	for i := 0; i < len(line); i++ {
		rn := line[i]
		val, exist := window[string(rn)]
		if !exist {
			window[string(rn)] = i
		} else {
			i = val + 1
			window = make(map[string]int, 4)
			window[string(line[i])] = i
		}
		if len(window) == length {
			fmt.Println(i+1, window)
			break
		}
	}
	// }
}
