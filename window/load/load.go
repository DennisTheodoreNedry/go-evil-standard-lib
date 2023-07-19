package load

import (
	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Sets an index file for usage
// This allow us to have more advance websites
func Load(index_file string, data_object *json.Json_t) {
	index_file = gotools.EraseDelimiter(index_file, []string{"\""}, -1)
	data_object.Add_index_file(index_file)
}
