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

func partOne(reports []string) {
	safeReports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		if isReportSafe(levels){
			safeReports++
		}
	}

	fmt.Println("Safe Reports:", safeReports)
}

func partTwo(reports []string) {
	safeReports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		for idx := range levels {
			newLevels := make([]string, len(levels) - 1)

			copy(newLevels, levels[:idx])
			copy(newLevels[idx:], levels[idx+1:])

			if isReportSafe(newLevels){
				safeReports++
				break
			}
		}
	}

	fmt.Println("Dampened Safe Reports:", safeReports)
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
	partTwo(reports)
}