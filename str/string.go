package str

import (
	"fmt"
	"strings"
)

func IsUpper(s string) (result bool) {
	return strings.ToUpper(s) == s
}

func IsLower(s string) (result bool) {
	return strings.ToLower(s) == s
}

func Capitalize(s string) (result string) {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func Snake(s string) (result string) {
	result = ""
	for index, ch := range s {
		if index == 0 {
			result += strings.ToLower(string(ch))
			continue
		}
		if IsUpper(string(ch)) {
			result += "_"
		}
		result += strings.ToLower(string(ch))
	}
	return result
}

func Dash(s string) (result string) {
	result = ""
	for index, ch := range s {
		if index == 0 {
			result += strings.ToLower(string(ch))
			continue
		}
		if IsUpper(string(ch)) {
			result += "-"
		}
		result += strings.ToLower(string(ch))
	}
	return result
}

func EscapeShellArg(s string) (result string) {
	backSlashEscapedStr := strings.ReplaceAll(s, "\\", "\\\\\\\\")
	quoteEscapedStr := strings.ReplaceAll(backSlashEscapedStr, "\"", "\\\"")
	backTickEscapedStr := strings.ReplaceAll(quoteEscapedStr, "`", "\\`")
	newLineEscapedStr := strings.ReplaceAll(backTickEscapedStr, "\n", "\\n")
	return fmt.Sprintf("\"%s\"", newLineEscapedStr)
}

func EscapeShellValue(s string, quote string) (result string) {
	backSlashEscapedStr := strings.ReplaceAll(s, "\\", "\\\\\\\\")
	quoteEscapedStr := backSlashEscapedStr
	if quote == "\"" {
		quoteEscapedStr = strings.ReplaceAll(backSlashEscapedStr, "\"", "\\\"")
	} else if quote == "'" {
		quoteEscapedStr = strings.ReplaceAll(backSlashEscapedStr, "'", "\\'")
	}
	backTickEscapedStr := strings.ReplaceAll(quoteEscapedStr, "`", "\\`")
	newLineEscapedStr := strings.ReplaceAll(backTickEscapedStr, "\n", "\\n")
	return newLineEscapedStr
}

func DoubleQuoteShellValue(s string) (result string) {
	return fmt.Sprintf("\"%s\"", EscapeShellValue(s, "\""))
}

func SingleQuoteShellValue(s string) (result string) {
	return fmt.Sprintf("'%s'", EscapeShellValue(s, "'"))
}

// GetSubKeys get sub keys from dictionary
func GetSubKeys(keys []string, parentKeys []string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for _, key := range keys {
		if !strings.HasPrefix(key, parentKey+"::") {
			continue
		}
		childKey := key[prefixLength:]
		if childKey == "" {
			continue
		}
		childKeyParts := strings.SplitN(childKey, "::", 2)
		subkey := childKeyParts[0]
		seen[subkey] = true
	}
	for key := range seen {
		subKeys = append(subKeys, key)
	}
	return subKeys
}

func GetUniqueElements(arr []string) (result []string) {
	result = []string{}
	seen := map[string]bool{}
	for _, element := range arr {
		if _, exist := seen[element]; exist {
			continue
		}
		result = append(result, element)
		seen[element] = true
	}
	return result
}

func Indent(multiLineStr string, indentation string) (indentedLine string) {
	lines := strings.Split(multiLineStr, "\n")
	return strings.Join(lines, "\n"+indentation)
}
