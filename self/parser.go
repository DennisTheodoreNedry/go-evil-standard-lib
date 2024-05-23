package main

import (
	"fmt"

	evil_call "github.com/DennisTheodoreNedry/go-evil/domains/self/call"
	"github.com/DennisTheodoreNedry/go-evil/domains/self/include"
	"github.com/DennisTheodoreNedry/go-evil/domains/self/random"
	"github.com/DennisTheodoreNedry/go-evil/domains/self/set"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "call":
		call = evil_call.Function(value, data_object)

	case "include":
		include.Include(value, data_object)

	case "set_run":
		call = set.Set(false, value, data_object)

	case "set_comp":
		call = set.Set(true, value, data_object)

	case "add_random_var":
		random.Add_variable(value, data_object)

	case "add_random_func":
		call = random.Add_function(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()", 1)

	}

	return call
}
