package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := parseInputFile("values.txt")
	if err != nil {
		fmt.Println("an error occurred while parsing the input file")
	}

	safeCount := 0
	unsafeCount := 0
	for _, report := range reports {
		if isSafeWithRemoval(report) {
			safeCount++
		} else {
			unsafeCount++
		}
	}

	fmt.Printf("Number of Safe Reports: %d\n", safeCount)
	fmt.Printf("Number of Unsafe Reports: %d\n", unsafeCount)
}

func checkSafety(report []int) bool {
	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if (report[i] > report[i-1]) != isIncreasing {
			return false
		}
	}

	return true
}

func isSafeWithRemoval(report []int) bool {
	if checkSafety(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if checkSafety(modifiedReport) {
			return true
		}
	}

	return false
}

func parseInputFile(filepath string) (reports [][]int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing the file {%s}: %v", filepath, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var report []int
		for _, num := range fields {
			value, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("failed to parse field : %q: %w", num, err)
			}
			report = append(report, value)
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("an error occurred while reading the file: %v", err)
	}

	return reports, nil
}
