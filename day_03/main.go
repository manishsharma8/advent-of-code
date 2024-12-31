package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const mulRegex = `(mul)\(\d{1,3},\d{1,3}\)`

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}

func partOne(input string) {
	var sum int

	mulRegexp, _ := regexp.Compile(mulRegex)
	matches := mulRegexp.FindAllString(input, -1)

	for _, match := range matches {
		values := strings.Split(match[4:len(match)-1], ",")
		sum += atoi(values[0]) * atoi(values[1])
	}

	fmt.Println("Multiplication Sum:", sum)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an input file")
		return;
	}

	file, err := os.ReadFile(os.Args[1])
	catch(err)

	partOne(string(file))
}
