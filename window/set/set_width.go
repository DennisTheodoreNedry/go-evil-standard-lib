package set

import (
	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Sets the width of the window
func Width(new_width string, data_object *json.Json_t) {
	new_width = gotools.EraseDelimiter(new_width, []string{"\""}, -1)

	data_object.Set_width(new_width)

}
