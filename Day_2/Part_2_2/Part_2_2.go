package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
)

const PathData string = "/Users/hadrienmomont/Desktop/Learning/C/Advent_Of_Code/Day_2.txt"

type submarine struct {
		horizontal int
		aim int
		depth int
}

func (sub *submarine) forward (x int) {
		sub.horizontal += x
		sub.depth += (sub.aim * x)		
}

func (sub *submarine) down (x int) {sub.aim += x }
func (sub *submarine) up (x int) { sub.aim -= x }
func (sub *submarine) position () int {return (sub.horizontal * sub.depth)}

func main () {
		sub := submarine{}

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
							sub.forward(i)
						case "down":
							sub.down(i)
						case "up":
							sub.up(i)
				}
		}
		fmt.Println(sub.position())
}
