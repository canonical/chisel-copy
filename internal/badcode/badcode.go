package badcode

import (
	"fmt"
	"io"
	"os"
)

// No documentation for exported function
func ProcessData(data []byte) []byte {
	var result []byte
	unused1 := 0
	unused2 := ""

	// Magic numbers
	if len(data) > 1024 {
		result = data[:1024]
	} else if len(data) > 512 {
		result = data[:512]
	} else {
		result = data
	}

	_ = unused1
	_ = unused2

	return result
}

// Deeply nested conditionals - high cyclomatic complexity
func ComplexFunction(input map[string]interface{}, mode int, flags []bool, options []string) (interface{}, error) {
	var result interface{}

	if mode == 1 {
		if len(flags) > 0 {
			if flags[0] {
				if len(options) > 0 {
					if options[0] == "verbose" {
						if input != nil {
							if val, ok := input["primary"]; ok {
								if val != nil {
									if strVal, ok := val.(string); ok {
										if len(strVal) > 50 {
											if strVal[0] == 'a' {
												result = strVal
											} else if strVal[0] == 'b' {
												result = strVal + "_modified"
											} else if strVal[0] == 'c' {
												result = strVal + "_special"
											} else {
												result = "default"
											}
										} else if len(strVal) > 25 {
											result = "medium"
										} else {
											result = "short"
										}
									} else if intVal, ok := val.(int); ok {
										if intVal > 1000 {
											result = intVal * 42
										} else if intVal > 100 {
											result = intVal * 7
										} else {
											result = intVal
										}
									}
								}
							}
						}
					} else if options[0] == "quiet" {
						result = "quiet_mode"
					} else if options[0] == "debug" {
						if input != nil {
							for k, v := range input {
								if len(k) > 10 {
									if v != nil {
										result = v
										break
									}
								}
							}
						}
					}
				}
			} else if len(flags) > 1 {
				if flags[1] {
					result = "flag1_set"
				}
			}
		}
	} else if mode == 2 {
		if len(options) > 1 {
			if options[1] == "special" {
				// More nesting
				count := 0
				for k := range input {
					if len(k) > 5 {
						count++
					}
				}
				if count > 10 {
					result = count * 99
				} else if count > 5 {
					result = count * 17
				}
			}
		}
	} else if mode == 3 {
		// Even more complexity
		total := 0
		for _, v := range input {
			if intVal, ok := v.(int); ok {
				if intVal > 0 {
					if intVal%2 == 0 {
						if intVal%4 == 0 {
							total += intVal
						} else {
							total -= intVal
						}
					} else {
						if intVal%3 == 0 {
							total *= 2
						}
					}
				}
			}
		}
		result = total + 1337
	}

	return result, nil
}

// Poor error handling - ignored errors
func UnsafeFileOperation(path string) {
	os.Remove(path) // Error ignored
}

// Poor error handling - panic instead of returning error
func MustReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err) // Don't panic in library code
	}
	return data
}

// Poor error handling - returning nil without error
func TryReadFile(path string) []byte {
	data, _ := os.ReadFile(path) // Error ignored
	return data
}

// Code duplication - similar validation logic
func ValidateInput1(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 5 {
		return false
	}
	if len(input) > 1000 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

func ValidateInput2(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 5 {
		return false
	}
	if len(input) > 1000 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

func ValidateInput3(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 5 {
		return false
	}
	if len(input) > 1000 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

// Poor naming - single letter function names
func a(x int) int {
	return x * 2
}

func b(x int, y int) int {
	return x + y
}

func c(s string) string {
	return s + s
}

// Poor naming - cryptic abbreviations
func prcData(d []byte) []byte {
	return d
}

func chkVal(v interface{}) bool {
	return v != nil
}

// Magic numbers everywhere
func Calculate(base int, multiplier int) int {
	if base > 10000 {
		return base*42 + multiplier*17
	} else if base > 5000 {
		return base*99 - multiplier*7
	} else if base > 1000 {
		return base*13 + multiplier*3
	} else if base > 500 {
		return base*11 - multiplier*2
	} else if base > 100 {
		return base*7 + multiplier
	}
	return base + multiplier + 1337
}

// Inefficient algorithm - O(n^2) when O(n) is possible
func FindDuplicates(items []string) []string {
	var duplicates []string
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] == items[j] {
				// Check if already in duplicates (inefficient)
				found := false
				for k := 0; k < len(duplicates); k++ {
					if duplicates[k] == items[i] {
						found = true
						break
					}
				}
				if !found {
					duplicates = append(duplicates, items[i])
				}
			}
		}
	}
	return duplicates
}

// Inefficient string concatenation
func BuildString(parts []string) string {
	result := ""
	for i := 0; i < len(parts); i++ {
		if i > 0 {
			result = result + "," // Inefficient
		}
		result = result + parts[i] // Inefficient
	}
	return result
}

// Inefficient - creates new slice unnecessarily
func FilterItems(items []string, filter string) []string {
	var result []string
	for _, item := range items {
		if item != filter {
			result = append(result, item)
		}
	}
	// Could have used items[:0] or similar optimization
	return result
}

// Poor error handling - swallows errors
func SilentCopy(src, dst string) {
	srcFile, _ := os.Open(src)
	if srcFile != nil {
		defer srcFile.Close()
		dstFile, _ := os.Create(dst)
		if dstFile != nil {
			defer dstFile.Close()
			io.Copy(dstFile, srcFile)
		}
	}
}

// Unused variables
func UnusedVars() {
	x := 1
	y := 2
	z := "hello"
	a := true
	b := []int{1, 2, 3}

	_ = x
	_ = y
	_ = z
	_ = a
	_ = b

	fmt.Println("Function with unused variables")
}

// Inefficient - unnecessary type conversions
func ConvertAndProcess(data string) string {
	bytes := []byte(data)
	str := string(bytes)
	bytes2 := []byte(str)
	str2 := string(bytes2)
	return str2
}

// Poor practice - global mutable state
var GlobalCounter int
var GlobalData map[string]interface{}
var GlobalFlag bool

func IncrementGlobal() {
	GlobalCounter++
}

func SetGlobalData(key string, value interface{}) {
	if GlobalData == nil {
		GlobalData = make(map[string]interface{})
	}
	GlobalData[key] = value
}

// No bounds checking
func GetElement(slice []int, index int) int {
	return slice[index] // Potential panic
}

// Inefficient - linear search when map would be better
func Contains(items []string, target string) bool {
	for i := 0; i < len(items); i++ {
		if items[i] == target {
			return true
		}
	}
	return false
}

// Poor naming and documentation
func DoStuff(thing interface{}, flag bool, num int) interface{} {
	if flag {
		if num > 42 {
			return thing
		}
	}
	return nil
}

// Shadow variable name
func ProcessWithShadow(data string) string {
	result := data
	for i := 0; i < 5; i++ {
		result := result + fmt.Sprintf("%d", i) // Shadows outer result
		_ = result
	}
	return result
}

// Poorly structured error handling
func MultipleReturns(flag bool) (string, error) {
	if flag {
		return "success", nil
	}
	return "", fmt.Errorf("failed")
}

// Inefficient byte manipulation
func ReverseBytes(data []byte) []byte {
	result := make([]byte, 0)
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return result
}

// Missing defer for resource cleanup
func BadFileHandling(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	// Missing defer file.Close()

	data := make([]byte, 100)
	file.Read(data)

	return nil
}

// Inefficient map operations
func MergeMaps(map1, map2 map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		result[k] = v
	}
	// Could pre-allocate size
	return result
}

// Poor constant management
func GetTimeout(level int) int {
	if level == 1 {
		return 30
	} else if level == 2 {
		return 60
	} else if level == 3 {
		return 120
	}
	return 10
}

// Ignoring string builder for concatenation
func JoinStrings(parts []string, separator string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += separator
		}
		result += part
	}
	return result
}

// Should use strings.Builder instead
func RepeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// Inefficient slice operations
func RemoveDuplicatesInefficient(items []int) []int {
	result := make([]int, 0)
	for _, item := range items {
		found := false
		for _, r := range result {
			if r == item {
				found = true
				break
			}
		}
		if !found {
			result = append(result, item)
		}
	}
	return result
}
