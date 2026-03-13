package statistics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputLines(path string) ([]string, error) {
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

func ParseInput(args []string) ([]int, []string, error) {
	nums := make([]int, 0, len(args))
	warnings := make([]string, 0)
	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("Skipping invalid number: %s", arg))
			continue
		}
		nums = append(nums, val)
	}

	if len(nums) == 0 {
		return nil, warnings, fmt.Errorf("no valid numbers provided")
	}

	return nums, warnings, nil
}
