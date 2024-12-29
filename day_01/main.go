package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func catch(e error) {
	if e != nil {
			panic(e)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}

func parseInput(lines []string) ([]int, []int) {
	var list1, list2 []int;

	for _, line := range lines {
		fields := strings.Fields(line)
		list1 = append(list1, atoi(fields[0]))
		list2 = append(list2, atoi(fields[1]))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2
}

func partOne(list1 []int, list2 []int) {
	var distance int = 0

	for idx := range list1 {
		distance += int(math.Abs(float64(list1[idx] - list2[idx])))
	}

	fmt.Println("Distance:", distance)
}

func partTwo(list1 []int, list2 []int) {
	similarityScore := 0
	freq := make(map[int]int)

	for _, i := range list2 {
		freq[i]++
	}

	for _, i := range list1 {
		similarityScore += i * freq[i];
	}

	fmt.Println("Similarity Score:", similarityScore)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an input file")
		return;
	}

	file, err := os.ReadFile(os.Args[1])
	catch(err)

	lines := strings.Split(string(file), "\n")
	list1, list2 := parseInput(lines)

	partOne(list1, list2)
	partTwo(list1, list2)
}