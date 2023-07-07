package generate

import (
	"fmt"

	tools "github.com/s9rA16Bf4/Go-tools"
)

func Generate_return_type() string {
	return_values := []string{"int", "bool", "string", "[]string", "[]int", "[]bool"}

	return return_values[tools.GenerateRandomIntBetween(0, len(return_values))]
}

func Generate_return_string_line(return_type string) string {
	to_return := ""

	switch return_type {
	case "int":
		to_return = fmt.Sprintf("return %d", tools.GenerateRandomInt())

	case "bool":
		to_return = fmt.Sprintf("return %t", tools.GenerateRandomBool())

	case "string":
		to_return = fmt.Sprintf("return \"%s\"", tools.Generate_random_string())

	case "[]bool":
		length := tools.GenerateRandomIntBetween(1, 64)
		line := "[]bool{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%t,", tools.GenerateRandomBool())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]int":
		length := tools.GenerateRandomIntBetween(1, 64)
		line := "[]int{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%d,", tools.GenerateRandomInt())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]string":
		length := tools.GenerateRandomIntBetween(1, 64)
		line := "[]string{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("\"%s\",", tools.Generate_random_string())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)
	}

	return to_return
}
