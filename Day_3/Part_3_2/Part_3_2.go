package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
//	"math"
)

const pathData string = "/Users/hadrienmomont/Desktop/Learning/Advent_Of_Code/Day_3.txt"

type diagnostic struct {
		sum_byte [12] float32
		key [12] byte
		oxygen string
		co2 [] byte
		len_input float32
}


// Method for the submarine struct. It basically needs the sum (per index) of all the input and look
func (s *diagnostic) find_key ()  {
	for index, value := range s.sum_byte {
		if common := value / s.len_input; float32(common) > 0.5 {
				s.key[index] = 1
		} else {
				s.key[index] = 0
			}
	}
}

func (s *diagnostic) oxygen_pattern (byte_input string) {
	check := false
	for index, value := range s.key {
		// Don't know how to do better than that yet. byte_input is a string, but when I call by index I get the integer value of that character
		str_to_int,_ := strconv.Atoi(string(byte_input[index])) 
		if int(value) == str_to_int {
			check = true
			//fmt.Println(value)
		} else {
			check = false
		}
	}
	if check {
		s.oxygen = byte_input
	}
}

func main () {
		file, err := os.Open(pathData)

		submarine := diagnostic {}

		if err != nil {
				panic(err)
		}

		scanner := bufio.NewScanner(file)
		

		for scanner.Scan() {
				buff := strings.Fields(scanner.Text())
				for _, byte_input := range buff {
					for i,j := range byte_input {
						a,_ := strconv.Atoi(string(j))
						submarine.sum_byte[i] += float32(a)
					}
				}
				submarine.len_input += 1
		}

		submarine.find_key()
		
		// Not the best but it seems that if I do not close the file, I can not use the scanner anymore. Perhpas it has something to do with a pointer, to investigate

		file.Close()
		file, err = os.Open(pathData)
		defer file.Close()

		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
				buff := strings.Fields(scanner.Text())
				for _, byte_input := range buff {
					//fmt.Println(byte_input)
					submarine.oxygen_pattern(byte_input)
				}
		}
			
		fmt.Println(submarine.oxygen)
		fmt.Println(submarine.key)



	}



