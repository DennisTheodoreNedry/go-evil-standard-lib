package io

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Reads the contents of a file and places the result into a runtime variable
func Read(value string, data_object *json.Json_t) []string {
	function_call := "read"
	value = gotools.EraseDelimiter(value, []string{"\""}, -1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1  []int"},
		Gut: []string{
			"path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"gut, err := ioutil.ReadFile(path)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"spine.variable.set(string(gut))",
		}})

	data_object.Add_go_import("io/ioutil")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}

}
