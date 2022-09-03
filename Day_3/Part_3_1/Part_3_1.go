package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

const pathData string = "/Users/hadrienmomont/Desktop/Learning/Advent_Of_Code/Day_3.txt"

type diagnostic struct {
		gamma [12] float32
		gamma_final [12] byte
		epsilon_final [12] byte
		len_input float32
}

// Simple function to reverse a byte array, it's for the espilon final byte array
func reverse (input [12] byte) [12] byte {
		var result [12] byte
		for i,j := range input {
			//fmt.Println(result[i])
			if j == 1 {
				result[i] = 0
			} else {
				result[i] = 1
			}
		}
		return result
}

// Function to convert binary number ( only works for this example thought

func binary_to_decimal (input [12] byte) float64 {
		lenght := len(input)
		var result float64

		for i,j := range(input) {
			result += float64(j) * math.Pow(2,float64(((lenght-1) - i)))
		}
		return result
}

// Method for the submarine struct. It basically needs the sum (per index) of all the input and look
func (s *diagnostic) process () uint {
	for index, value := range s.gamma {
		if common := value / s.len_input; float32(common) > 0.5 {
				s.gamma_final[index] = 1
		} else {
				s.gamma_final[index] = 0
			}
	}

	s.epsilon_final = reverse(s.gamma_final)
	gamma_decimal := binary_to_decimal(s.gamma_final)
	epsilon_final := binary_to_decimal(s.epsilon_final)

	power_consumption := epsilon_final * gamma_decimal
	return uint(power_consumption)
}


func main () {
		file, err := os.Open(pathData)
		defer file.Close()	

		submarine := diagnostic {}

		if err != nil {
				panic(err)
		}

		scanner := bufio.NewScanner(file)
		
		fmt.Println(scanner.Text())

		for scanner.Scan() {

			// Simple, I read by line the dataset and then I loop through the line
			// Then I keep on summing each value of each index to get a .gamma that is the bug sum. 
			// Then I the method I divide by the lenght of the dataset and therefore I get a final byte array for the gamma
			// I just have to reverse it for the epsilon, covert the value and multiply and it is over
				buff := strings.Fields(scanner.Text())
				for _, byte_input := range buff {
					for i,j := range byte_input {
						a,_ := strconv.Atoi(string(j))
						submarine.gamma[i] += float32(a)
					}
				}
				// Just a counter to get the lenght of the dataset so it is store to the struct
				submarine.len_input += 1
		}

		power := submarine.process()
		fmt.Println(power)




}

