package main

import (
	"fmt"
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

func partOne(grid [][]rune) {
	count := 0
	dirs := []int{-1, 0, 1}

	for r, row := range grid {
		for c, cell := range row {
			if cell != 'X' {
				continue
			}

			for _, dr := range dirs {
				for _, dc := range dirs {
					if dr == 0 && dc == 0 {
						continue
					}

					nr, nc := r + 3 * dr, c + 3 * dc
					if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(row) {
						continue
					}

					if grid[r + dr][c + dc] == 'M' && grid[r + 2 * dr][c + 2 * dc] == 'A' && grid[nr][nc] == 'S' {
						count++
					}
				}
			}
		}
	}

	fmt.Println("XMAS Count:", count)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an input file")
		return;
	}

	file, err := os.ReadFile(os.Args[1])
	catch(err)

	grid := strings.Split(string(file), "\n")
	runeGrid := make([][]rune, len(grid))

	for i, row := range grid {
		runeGrid[i] = []rune(row)
	}

	partOne(runeGrid)
}
