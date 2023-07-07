package main

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/infect/path"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "path":
		call = path.Path(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "infect.Parser()", 1)

	}

	return call
}
