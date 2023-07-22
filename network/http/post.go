package http

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Peforms a POST Request to the provided url
// The input must be an evil array with the format ${"url", "data", "agent", "content-type"}$
// The result is saved in a runtime variable
func Post(value string, data_object *json.Json_t) []string {
	function_call := "post"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 4 {
		notify.Error(fmt.Sprintf("Expected four values, but recieved %d", arr.Length()), "http.Post()", 1)
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 []int", "repr_3 []int", "repr_4 []int"},
		Gut: []string{
			"url := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"data := []byte(spine.variable.get(spine.alpha.construct_string(repr_2)))",
			"agent := spine.variable.get(spine.alpha.construct_string(repr_3))",
			"content_type := spine.variable.get(spine.alpha.construct_string(repr_4))",

			"data_reader := bytes.NewReader(data)",

			"req, err := http.NewRequest(http.MethodPost, url, data_reader)",

			"if err != nil {",
			"spine.log(err.Error())",
			"}",

			"req.Header.Set(\"Content-Type\", content_type)",
			"req.Header.Set(\"User-Agent\", agent)",
			"client := &http.Client{}",
			"resp, err := client.Do(req)",

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
	data_object.Add_go_import("bytes")

	parameter_1 := data_object.Generate_int_array_parameter(arr.Get(0))
	parameter_2 := data_object.Generate_int_array_parameter(arr.Get(1))
	parameter_3 := data_object.Generate_int_array_parameter(arr.Get(2))
	parameter_4 := data_object.Generate_int_array_parameter(arr.Get(3))

	return []string{fmt.Sprintf("%s(%s, %s, %s, %s)", function_call, parameter_1, parameter_2, parameter_3, parameter_4)}
}
