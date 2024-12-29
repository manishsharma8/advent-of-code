package main

import (
	"fmt"
	"math"
	"os"
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

func isReportSafe(levels []string) bool {
	initialDifference := atoi(levels[1]) - atoi(levels[0])
	isMonotonic := initialDifference > 0

	for i := 1; i < len(levels); i++ {
		currentDifference := atoi(levels[i]) - atoi(levels[i-1])

		if (currentDifference > 0) != isMonotonic {
			return false
		}

		if math.Abs(float64(currentDifference)) > 3 || math.Abs(float64(currentDifference)) < 1 {
			return false
		}
	}

	return true
}

func partOne(report []string) {
	safeReports := 0

	for _, report := range report {
		levels := strings.Split(report, " ")

		if isReportSafe(levels){
			safeReports++
		}
	}

	fmt.Println("Safe Reports:", safeReports)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an input file")
		return;
	}

	file, err := os.ReadFile(os.Args[1])
	catch(err)

	reports := strings.Split(string(file), "\n")
	partOne(reports)
}