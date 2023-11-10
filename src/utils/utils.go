package utils

import (
	"bufio"
	"fmt"
	"os"
)

// read line by line into memory
// all file contents is stores in lines[]
func ReadLines(path string) []string {
	path = "../" + path
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ToDo how can I make this polymorphic for ints and strings?
func Check_answer(ans int, expect int) bool {
	if ans == expect {
		fmt.Println("Test Passed")
		return true
	} else {
		fmt.Println("Test Failed, expected:", expect, "received:", ans)
		suite_passed = false
		return false
	}
}

func Check_all() bool {
	if suite_passed {
		println("\nAll tests passed!")
		return true
	} else {
		println("\nSome tests failed")
		return false
	}
}

var suite_passed = true
