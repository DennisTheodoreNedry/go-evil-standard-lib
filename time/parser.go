package main

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/domains/time/sleep"
	"github.com/DennisTheodoreNedry/go-evil/domains/time/until"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "until":
		call = until.Until(value, data_object)

	case "sleep":
		call = sleep.Sleep(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()", 1)

	}

	return call
}
