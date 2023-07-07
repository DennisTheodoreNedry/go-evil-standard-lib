package main

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/base64/decode"
	"github.com/s9rA16Bf4/go-evil/domains/base64/encode"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
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
