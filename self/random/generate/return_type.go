package generate

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
)

func Generate_return_type() string {
	return_values := []string{"int", "bool", "string", "[]string", "[]int", "[]bool"}

	return return_values[gotools.GenerateRandomIntBetween(0, len(return_values))]
}

func Generate_return_string_line(return_type string) string {
	to_return := ""

	switch return_type {
	case "int":
		to_return = fmt.Sprintf("return %d", gotools.GenerateRandomInt())

	case "bool":
		to_return = fmt.Sprintf("return %t", gotools.GenerateRandomBool())

	case "string":
		to_return = fmt.Sprintf("return \"%s\"", gotools.Generate_random_string())

	case "[]bool":
		length := gotools.GenerateRandomIntBetween(1, 64)
		line := "[]bool{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%t,", gotools.GenerateRandomBool())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]int":
		length := gotools.GenerateRandomIntBetween(1, 64)
		line := "[]int{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("%d,", gotools.GenerateRandomInt())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)

	case "[]string":
		length := gotools.GenerateRandomIntBetween(1, 64)
		line := "[]string{"
		for z := 0; z < length; z++ {
			line += fmt.Sprintf("\"%s\",", gotools.Generate_random_string())
		}
		line += "}"

		to_return = fmt.Sprintf("return %s", line)
	}

	return to_return
}
