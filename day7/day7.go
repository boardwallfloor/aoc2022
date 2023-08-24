package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	types string
	name  string
	item  []*Node
	size  int
}

func getIndent(depth int) string {
	return strings.Repeat(" ", depth)
}

func TraverseNode(node *Node, depth int, dir *[]*Node) (int, int) {
	tot := node.size
	mark := 0
	if node == nil {
		return tot, mark
	}

	for _, br := range node.item {
		total, track := TraverseNode(br, depth+1, dir)
		tot += total
		mark += track
	}

	if node.types == "dir" {
		*dir = append(*dir, node)
		node.size = tot
	}
	// fmt.Printf("%s-%s, %d, %v, %d\n", getIndent(depth), node.name, node.size, node.item, tot)

	return tot, mark
}

func main() {
	// 	input := `$ cd /
	// $ ls
	// dir a
	// 14848514 b.txt
	// 8504156 c.dat
	// dir d
	// $ cd a
	// $ ls
	// dir e
	// 29116 f
	// 2557 g
	// 62596 h.lst
	// $ cd e
	// $ ls
	// 584 i
	// $ cd ..
	// $ cd ..
	// $ cd d
	// $ ls
	// 4060174 j
	// 8033020 d.log
	// 5626152 d.ext
	// 7214296 k`

	input, err := os.Open("./day7")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	// lines := strings.Split(input, "\n")
	history := make([]*Node, 0)
	paths := make([]string, 0)
	isList := false
	var cwd *Node
	// for _, line := range lines {
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == byte('$') {
			isList = false

			if line[2] == byte('c') {
				path := line[5:]

				if line[5:] == "/" {
					cwd = &Node{name: path, types: "dir"}
					paths = append(paths, path)
					history = append(history, cwd)
				} else if line[5:] != ".." {
					for _, child := range cwd.item {
						if child.name == path {
							cwd = child
							break
						}
					}
					paths = append(paths, path, "/")
					history = append(history, cwd)
				}
				if line[5:] == ".." {
					history = history[:len(history)-1]
					cwd = history[len(history)-1]
				}
			}

			if line[2] == byte('l') {
				isList = true
			}

		} else {
			if isList {
				file := strings.Split(line, " ")
				if line[0] == byte('d') {
					item := &Node{name: file[1], types: "dir"}
					cwd.item = append(cwd.item, item)
				} else {
					fileSize, err := strconv.Atoi(file[0])
					if err != nil {
						fmt.Println(err, line)
					}
					item := &Node{name: file[1], size: fileSize, types: "file"}
					cwd.item = append(cwd.item, item)
				}

			}
		}
	}
	dir := make([]*Node, 0)
	total, _ := TraverseNode(history[0], 0, &dir)
	free := 70000000 - total
	diff := 30000000 - free
	temp := make([]int, 0)
	for _, v := range dir {
		if v.size >= diff {
			temp = append(temp, v.size)
		}
	}
	slices.Sort(temp)
	fmt.Println(temp)
}
