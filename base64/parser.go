package main

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/domains/base64/decode"
	"github.com/DennisTheodoreNedry/go-evil/domains/base64/encode"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "encode":
		call = encode.Encode(value, data_object)

	case "decode":
		call = decode.Decode(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "base64.Parser()", 1)

	}

	return call
}
