package generate

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Generates a rsa key used for encrypting/decrypting
func RSA_key(value string, data_object *json.Json_t) []string {
	value = gotools.EraseDelimiter(value, []string{"\""}, -1)
	function_call := "generate_rsa_key"

	// Check if the key is valid
	if ok := gotools.StringToInt(value); ok == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()", 1)
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"key_size int"},
		Gut: []string{
			"privateKey, err := rsa.GenerateKey(rand.Reader, key_size)",
			"if err == nil{",
			"spine.crypt.set_rsa_key(privateKey, key_size)",
			"}",
		}})

	return []string{fmt.Sprintf("%s(%s)", function_call, value)}
}
