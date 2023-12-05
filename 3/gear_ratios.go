package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	//"strings"
	"strconv"
)

var debug bool = false
var symbols = []string{
	"!",
	"/",
	"%",
	"$",
	"=",
	"+",
	"*",
	"@",
	"-",
	"#",
	"&",
}

func main() {

	var input_file string = "input.txt"

	if debug {
		input_file = "debug_input.txt"
	}

	content, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	inputText := parseInputFile(scanner)

	sum := getNumbers(inputText)
	fmt.Println(sum)

}

func parseInputFile(scanner *bufio.Scanner) [][]string {
	inputText := [][]string{}
	for scanner.Scan() || len(scanner.Text()) > 0 {
		txt := scanner.Text()
		arr := []string{}
		for _, s := range txt {
			arr = append(arr, string(s))
		}
		inputText = append(inputText, arr)
	}
	if debug {
		fmt.Println(inputText)
	}
	return inputText

}

func getNumbers(inputText [][]string) int {
	var sum int = 0
	var i, j int = 0, 0
	for i < len(inputText) {
		for j < len(inputText[i]) {
			if isNumber(inputText[i][j]) {
				num := ""
				k := j
				for j < len(inputText[i]) && isNumber(inputText[i][j]) {
					num += inputText[i][j]
					j++
				}
				if isThereSymbolNearBy(inputText, i, k, j) {
					if intNum, numOk := strconv.Atoi(num); numOk == nil {
						fmt.Println(num)
						sum += intNum
						//fmt.Println(sum)
					}
				}
			} else if j >= len(inputText[i]) {
				break
			} else {
				j++
			}
		}
		if j == len(inputText[i]) {
			fmt.Printf("/\n")
			i++
			j = 0
		}
	}
	return sum
}

func isThereSymbolNearBy(inputText [][]string, i int, left int, right int) bool {

	rowStart := i - 1
	rowEnd := i + 1

	columnStart := left - 1
	columnEnd := right

	// check the array bounds
	if rowStart < 0 {
		rowStart = i
	}
	if rowEnd >= len(inputText) {
		rowEnd = i
	}
	if columnStart < 0 {
		columnStart = left
	}
	if columnEnd >= len(inputText[i]) {
		columnEnd = right - 1
	}
	// check if there is symbol near by
	for row := rowStart; row <= rowEnd; row++ {
		for column := columnStart; column <= columnEnd; column++ {
			symbol := inputText[row][column]
			if stringInSlice(symbol, symbols) {
				//fmt.Println("There is a symbol nearby")
				return true
			}
		}
	}
	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func isStar(a string) bool {
    return a == "*"
}
