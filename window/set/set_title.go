package set

import (
	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Sets the title of the window that appears
func Title(new_title string, data_object *json.Json_t) {
	new_title = gotools.EraseDelimiter(new_title, []string{"\""}, -1)

	data_object.Set_title(new_title)

}
