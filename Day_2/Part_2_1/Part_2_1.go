package day_2_1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

const PathData string = "/Users/hadrienmomont/Desktop/Learning/C/Advent_Of_Code/Day_2.txt"

func (s *Submarine_pos) Position () int {return (s.horizontal * s.depth)}

type Submarine_pos struct {
		Horizontal int
		Depth int
	}

func day_2_1 () {
		sub := Submarine_pos{}

		file, err := os.Open(PathData)
		defer file.Close()	

		if err != nil {
				panic(err)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
				b := strings.Fields(scanner.Text())

				switch i,_ := strconv.Atoi(b[1]); b[0] {
						case "forward":
							sub.Horizontal += i
						case "down":
							sub.Depth += i
						case "up":
							sub.Depth -= i
				}

		}

		fmt.Printf("Answer:\t%d\n", sub.Position())
}
