package testutil

import (
	"fmt"
	"os"
	"strings"
)

// UntestedFunction1 has no corresponding test
func UntestedFunction1(input string) string {
	if input == "" {
		return "empty"
	}
	return strings.ToUpper(input)
}

// UntestedFunction2 has no test coverage
func UntestedFunction2(a, b int) int {
	if a > b {
		return a - b
	} else if a < b {
		return b - a
	}
	return 0
}

// UntestedComplexFunction has no test and high cyclomatic complexity
func UntestedComplexFunction(data map[string]interface{}, mode int) (interface{}, error) {
	if mode == 1 {
		if val, ok := data["key1"]; ok {
			if strVal, ok := val.(string); ok {
				if len(strVal) > 10 {
					return strVal, nil
				} else if len(strVal) > 5 {
					return strVal + "_short", nil
				} else {
					return "very_short", nil
				}
			}
		}
	} else if mode == 2 {
		for k, v := range data {
			if k == "special" {
				return v, nil
			}
		}
	} else if mode == 3 {
		// Magic numbers
		count := 0
		for range data {
			count++
			if count > 42 {
				break
			}
		}
		return count * 17, nil
	}
	return nil, fmt.Errorf("invalid mode: %d", mode)
}

// UntestedFileOperation performs file operations without tests
func UntestedFileOperation(path string) error {
	data, _ := os.ReadFile(path) // Error ignored
	if len(data) > 0 {
		os.WriteFile(path+".bak", data, 0644) // Error ignored
	}
	return nil
}

// UntestedDataProcessor processes data inefficiently
func UntestedDataProcessor(items []string) []string {
	var result []string
	// Inefficient: O(n^2) when O(n) is possible
	for i := 0; i < len(items); i++ {
		duplicate := false
		for j := 0; j < len(result); j++ {
			if result[j] == items[i] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, items[i])
		}
	}
	return result
}

// x - poorly named function with no test
func x(y int) int {
	return y * 2
}

// z - another poorly named function with no test
func z(s string) bool {
	return len(s) > 5
}

// UntestedPanicFunction panics instead of returning errors
func UntestedPanicFunction(input string) string {
	if input == "" {
		panic("empty input not allowed")
	}
	return input
}

// UntestedDuplicateLogic1 has duplicated code
func UntestedDuplicateLogic1(path string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) < 3 {
		return false
	}
	if len(path) > 200 {
		return false
	}
	for i := 0; i < len(path); i++ {
		if path[i] < 32 {
			return false
		}
	}
	return true
}

// UntestedDuplicateLogic2 has nearly identical code to UntestedDuplicateLogic1
func UntestedDuplicateLogic2(path string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) < 3 {
		return false
	}
	if len(path) > 200 {
		return false
	}
	for i := 0; i < len(path); i++ {
		if path[i] < 32 {
			return false
		}
	}
	return true
}
