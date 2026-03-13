package statistics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadDataLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseInput(lines []string) ([]int, []string, error) {
	nums := make([]int, 0, len(lines))
	parseWarnings := make([]string, 0)
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			parseWarnings = append(parseWarnings, fmt.Sprintf("Skipping invalid number: %s", line))
			continue
		}
		nums = append(nums, val)
	}

	if len(nums) == 0 {
		return nil, parseWarnings, fmt.Errorf("no valid numbers provided")
	}

	return nums, parseWarnings, nil
}
