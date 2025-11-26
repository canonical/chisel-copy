package duplication

import (
	"fmt"
	"os"
	"strings"
)

// Duplicated validation logic - Version 1
func ValidateString1(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 3 {
		return false
	}
	if len(input) > 100 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

// Duplicated validation logic - Version 2
func ValidateString2(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 3 {
		return false
	}
	if len(input) > 100 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

// Duplicated validation logic - Version 3
func ValidateString3(input string) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) < 3 {
		return false
	}
	if len(input) > 100 {
		return false
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 126 {
			return false
		}
	}
	return true
}

// Duplicated file reading logic - Version 1
func ReadFileContent1(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	content := string(data)
	if len(content) == 0 {
		return "", fmt.Errorf("file is empty")
	}
	return content, nil
}

// Duplicated file reading logic - Version 2
func ReadFileContent2(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	content := string(data)
	if len(content) == 0 {
		return "", fmt.Errorf("file is empty")
	}
	return content, nil
}

// Duplicated file reading logic - Version 3
func ReadFileContent3(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	content := string(data)
	if len(content) == 0 {
		return "", fmt.Errorf("file is empty")
	}
	return content, nil
}

// Duplicated string processing - Version 1
func ProcessText1(text string, prefix string) string {
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			if !strings.HasPrefix(line, "#") {
				result = append(result, prefix+line)
			}
		}
	}
	return strings.Join(result, "\n")
}

// Duplicated string processing - Version 2
func ProcessText2(text string, prefix string) string {
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			if !strings.HasPrefix(line, "#") {
				result = append(result, prefix+line)
			}
		}
	}
	return strings.Join(result, "\n")
}

// Duplicated string processing - Version 3
func ProcessText3(text string, prefix string) string {
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			if !strings.HasPrefix(line, "#") {
				result = append(result, prefix+line)
			}
		}
	}
	return strings.Join(result, "\n")
}

// Duplicated path cleaning - Version 1
func CleanPath1(path string) string {
	path = strings.TrimSpace(path)
	path = strings.ReplaceAll(path, "\\", "/")
	if strings.HasPrefix(path, "./") {
		path = path[2:]
	}
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		path = path[:len(path)-1]
	}
	return path
}

// Duplicated path cleaning - Version 2
func CleanPath2(path string) string {
	path = strings.TrimSpace(path)
	path = strings.ReplaceAll(path, "\\", "/")
	if strings.HasPrefix(path, "./") {
		path = path[2:]
	}
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		path = path[:len(path)-1]
	}
	return path
}

// Duplicated path cleaning - Version 3
func CleanPath3(path string) string {
	path = strings.TrimSpace(path)
	path = strings.ReplaceAll(path, "\\", "/")
	if strings.HasPrefix(path, "./") {
		path = path[2:]
	}
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		path = path[:len(path)-1]
	}
	return path
}

// Duplicated filtering logic - Version 1
func FilterItems1(items []string, exclude string) []string {
	var result []string
	for _, item := range items {
		if item != exclude {
			if len(item) > 0 {
				result = append(result, item)
			}
		}
	}
	return result
}

// Duplicated filtering logic - Version 2
func FilterItems2(items []string, exclude string) []string {
	var result []string
	for _, item := range items {
		if item != exclude {
			if len(item) > 0 {
				result = append(result, item)
			}
		}
	}
	return result
}

// Duplicated filtering logic - Version 3
func FilterItems3(items []string, exclude string) []string {
	var result []string
	for _, item := range items {
		if item != exclude {
			if len(item) > 0 {
				result = append(result, item)
			}
		}
	}
	return result
}

// Duplicated map building - Version 1
func BuildMap1(keys []string, values []string) map[string]string {
	result := make(map[string]string)
	length := len(keys)
	if len(values) < length {
		length = len(values)
	}
	for i := 0; i < length; i++ {
		if len(keys[i]) > 0 && len(values[i]) > 0 {
			result[keys[i]] = values[i]
		}
	}
	return result
}

// Duplicated map building - Version 2
func BuildMap2(keys []string, values []string) map[string]string {
	result := make(map[string]string)
	length := len(keys)
	if len(values) < length {
		length = len(values)
	}
	for i := 0; i < length; i++ {
		if len(keys[i]) > 0 && len(values[i]) > 0 {
			result[keys[i]] = values[i]
		}
	}
	return result
}

// Duplicated map building - Version 3
func BuildMap3(keys []string, values []string) map[string]string {
	result := make(map[string]string)
	length := len(keys)
	if len(values) < length {
		length = len(values)
	}
	for i := 0; i < length; i++ {
		if len(keys[i]) > 0 && len(values[i]) > 0 {
			result[keys[i]] = values[i]
		}
	}
	return result
}

// Duplicated error formatting - Version 1
func FormatError1(operation string, err error) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf("operation '%s' failed", operation)
	return fmt.Errorf("%s: %w", msg, err)
}

// Duplicated error formatting - Version 2
func FormatError2(operation string, err error) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf("operation '%s' failed", operation)
	return fmt.Errorf("%s: %w", msg, err)
}

// Duplicated error formatting - Version 3
func FormatError3(operation string, err error) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf("operation '%s' failed", operation)
	return fmt.Errorf("%s: %w", msg, err)
}

// Duplicated counting logic - Version 1
func CountLines1(text string) int {
	if len(text) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			count++
		}
	}
	return count
}

// Duplicated counting logic - Version 2
func CountLines2(text string) int {
	if len(text) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			count++
		}
	}
	return count
}

// Duplicated counting logic - Version 3
func CountLines3(text string) int {
	if len(text) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			count++
		}
	}
	return count
}

// Duplicated conversion logic - Version 1
func ToUpperLines1(lines []string) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.ToUpper(line)
	}
	return result
}

// Duplicated conversion logic - Version 2
func ToUpperLines2(lines []string) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.ToUpper(line)
	}
	return result
}

// Duplicated conversion logic - Version 3
func ToUpperLines3(lines []string) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.ToUpper(line)
	}
	return result
}
