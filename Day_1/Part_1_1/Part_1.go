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

	for i:= 1 ; i <= len(data)-1; i++ {
		if (*p)[i] >= (*p)[i-1] {
			counter += 1
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
