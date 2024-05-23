package random

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Adds a random variable to the source code
func Add_variable(amount string, data_object *json.Json_t) {

	amount = gotools.EraseDelimiter(amount, []string{"\""}, -1)

	i_value := gotools.StringToInt(amount)
	if i_value == -1 {
		notify.Error(fmt.Sprintf("Unknown amount '%d'", i_value), "self.Add_random_variable()", 1)
	}

	for i := 0; i < i_value; i++ {
		variable_name := gotools.Generate_random_string()
		random_value := gotools.Generate_random_string()

		data_object.Add_go_global(fmt.Sprintf("var %s string = \"%s\"", variable_name, random_value))
	}
}
