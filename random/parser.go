package main

import (
	"fmt"

	evil_int "github.com/DennisTheodoreNedry/go-evil/domains/random/int"
	evil_string "github.com/DennisTheodoreNedry/go-evil/domains/random/string"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "int":
		evil_int.Generate(value, data_object)

	case "string":
		evil_string.Generate(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()", 1)

	}

	return call
}
