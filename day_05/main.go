package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func parseRule(line string) (int, int) {
	parts := strings.Split(line, "|")
	before, _ := strconv.Atoi(parts[0])
	after, _ := strconv.Atoi(parts[1])
	return before, after
}

func parsePacket(line string) []int {
	packet := make([]int, 0)
	for _, value := range strings.Split(line, ",") {
		ival, _ := strconv.Atoi(value)
		packet = append(packet, ival)
	}
	return packet
}

func parseInput(input string) (map[[2]int]bool, [][]int) {
	lines := strings.Split(input, "\n")

	rules := make(map[[2]int]bool)
	packets := make([][]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			before, after := parseRule(line)
			rules[[2]int{before, after}] = true
			rules[[2]int{after, before}] = false
		} else {
			packet := parsePacket(line)
			packets = append(packets, packet)
		}
	}

	return rules, packets
}

func isOrdered(packet []int, rules map[[2]int]bool) bool {
	for i, value1 := range packet {
		for _, value2 := range packet[i + 1:] {
			key, exist := rules[[2]int{value1, value2}]
			if exist && !key {
				return false
			}
		}
	}
	return true;
}

func partOne(rules map[[2]int]bool, packets [][]int) {
	middleSum := 0

	for _, packet := range packets {
		if isOrdered(packet, rules) {
			middleSum += packet[len(packet) / 2]
		}
	}

	fmt.Println("Sum (ordered updates)", middleSum)
}

func partTwo(rules map[[2]int]bool, packets [][]int) {
	middleSum := 0

	for _, packet := range packets {
		if !isOrdered(packet, rules) {
			sort.Slice(packet, func(i, j int) bool {
				key, exist := rules[[2]int{packet[i], packet[j]}]
				return exist && !key
			})
			middleSum += packet[len(packet) / 2]
		}
	}

	fmt.Println("Sum (incorrectly-ordered updates)", middleSum)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an input file")
		return;
	}

	file, err := os.ReadFile(os.Args[1])
	catch(err)

	rules, packets := parseInput(string(file))

	partOne(rules, packets)
	partTwo(rules, packets)
}
