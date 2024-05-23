package string

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Generates a random string where the provded value is the length of the string
// The generated value is placed in a compile-time variable
func Generate(value string, data_object *json.Json_t) {

	roof := gotools.EraseDelimiter(value, []string{"\""}, -1)
	length := gotools.StringToInt(roof)

	if length == -1 {
		notify.Error(fmt.Sprintf("Failed to convert value %s to an integer!", roof), "random.generate_string()", 1)
	}

	generated_value := gotools.Generate_random_n_string(length)

	data_object.Set_variable_value(generated_value)

}
