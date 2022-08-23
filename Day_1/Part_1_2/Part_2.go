package main

import (
	"fmt"
	"os"
	"strconv"
)

const pathData string = "/Users/hadrienmomont/Desktop/Learning/C/Advent_Of_Code/Part_1.txt"


func main () {
	file, err := os.ReadFile(pathData)

	if err != nil {
			panic(err)
	}

	data := split(file)
	
	var counter int = 0
	p := &data
	k := 1
	for i := 0 ; i <= len(data) - 4; i, k = i + 1, k + 1 {

		buff_0 := (*p)[i] + (*p)[i+1] + (*p)[i+2]
		buff_1 := (*p)[k] + (*p)[k+1] + (*p)[k+2]

		if (buff_1 - buff_0 > 0) {
			counter++
		}
	}
	
	fmt.Println(counter)
}

// Just a small training and making my own split function. 
func split (input [] byte) [] int {

		var result [] int
		var flag int
		
		for j, element := range input {

			if  element == '\n' || element == '\r' || element == '\t' {	
					num, _ := strconv.Atoi(string(input[flag:j]))
					result = append(result,num)
					flag = j+1
				}
		}
		return result
}
