package set

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Sets the height of the window
func Height(new_height string, data_object *json.Json_t) {
	new_height = gotools.EraseDelimiter(new_height, []string{"\""}, -1)

	data_object.Set_height(new_height)

}
