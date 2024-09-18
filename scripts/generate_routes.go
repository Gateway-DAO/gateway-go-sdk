package scripts

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func GenerateRouteConstants(paths map[string]PathItem) string {
	var constants []string
	constants = append(constants, "const (")

	for path, pathItem := range paths {
		for method, operation := range pathItem {
			constantName := generateRouteConstantName(method, path, operation.Summary)
			parameterizedPath := parameterizePath(path)
			constants = append(constants, fmt.Sprintf("\t%s = \"%s\"", constantName, parameterizedPath))
		}
	}

	constants = append(constants, ")")
	return strings.Join(constants, "\n")
}

func generateRouteConstantName(method, path, summary string) string {
	name := strings.TrimSpace(summary)
	if name == "" {
		name = path
	}
	return toPascalCase(name)
}

func parameterizePath(path string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	return re.ReplaceAllString(path, "%s")
}

func toPascalCase(s string) string {
	var sb strings.Builder
	uppercaseNext := true

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			uppercaseNext = true
		} else {
			if uppercaseNext {
				sb.WriteRune(unicode.ToUpper(r))
			} else {
				sb.WriteRune(r)
			}
			uppercaseNext = false
		}
	}

	return sb.String()
}
