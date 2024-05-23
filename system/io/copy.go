package io

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Copies the target file to the new provided location
func Copy(value string, data_object *json.Json_t) []string {
	function_call := "copy"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()", 1)
	}

	old_path := arr.Get(0)
	new_path := arr.Get(1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 []int"},
		Gut: []string{
			"old_path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"new_path := spine.variable.get(spine.alpha.construct_string(repr_2))",

			"src, err := os.Open(old_path)",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",

			"dst, err := os.Create(new_path)",
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",

			"_, err = io.Copy(dst, src)",

			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("io")
	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/DennisTheodoreNedry/notify_handler")

	// Construct our int array
	old_parameter := data_object.Generate_int_array_parameter(old_path)
	new_parameter := data_object.Generate_int_array_parameter(new_path)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, old_parameter, new_parameter)}
}
