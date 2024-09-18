package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SwaggerSchema struct {
	Definitions map[string]struct {
		Type          string                 `json:"type"`
		Properties    map[string]SwaggerType `json:"properties"`
		Enum          []string               `json:"enum"`
		XEnumVarnames []string               `json:"x-enum-varnames"`
	} `json:"definitions"`
	Paths map[string]PathItem `json:"paths"`
}

type PathItem map[string]Operation

type Operation struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type SwaggerType struct {
	Type string `json:"type"`
	Ref  string `json:"$ref"`
}

func GenerateTypes() {
	data, err := os.ReadFile("api.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var schema SwaggerSchema
	err = json.Unmarshal(data, &schema)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	goTypes := extractGoTypes(schema)

	routeConstants := GenerateRouteConstants(schema.Paths)

	outputTypes := "package common\n\n" + strings.Join(goTypes, "\n\n") + "\n\n"
	err = os.WriteFile("pkg/common/types.go", []byte(outputTypes), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	outputRoutesConstants := "package common\n\n" + routeConstants + "\n\n"
	err = os.WriteFile("pkg/common/routes.go", []byte(outputRoutesConstants), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Go types have been generated and saved to files.")

}
