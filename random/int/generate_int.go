package int

import (
	"fmt"

	tools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Generates a random int value in a range
// The input is an evil array that should only contain two values, min and max
// The generated value is placed in a compile-time variable
func Generate(value string, data_object *json.Json_t) {
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected an array with two positions, got %d", arr.Length()), "random.generate_int()", 1)
	}

	min := arr.Get(0)
	i_min := tools.StringToInt(min)

	if i_min == -1 {
		notify.Error(fmt.Sprintf("Failed to convert min value %s to an integer!", min), "random.generate_int()", 1)
	}

	max := arr.Get(1)
	i_max := tools.StringToInt(max)

	if i_max == -1 {
		notify.Error(fmt.Sprintf("Failed to convert max value %s to an integer!", max), "random.generate_int()", 1)
	}

	generated_value := tools.GenerateRandomIntBetween(i_min, i_max)

	data_object.Set_variable_value(tools.IntToString(generated_value))

}
