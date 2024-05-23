package generate

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
)

func Generate_parameters() ([]string, string) {
	amount_of_parameters := gotools.GenerateRandomInt()
	sending_values := ""
	parameter_array := []string{}

	for y := 0; y < amount_of_parameters; y++ {
		parameter := fmt.Sprintf("param%d string", y)
		parameter_array = append(parameter_array, parameter)

		sending_values += fmt.Sprintf("\"%s\",", gotools.Generate_random_string())
	}

	return parameter_array, sending_values
}
