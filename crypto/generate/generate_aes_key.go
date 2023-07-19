package generate

import (
	"fmt"

	gotools "github.com/s9rA16Bf4/Go-tools"
	"github.com/s9rA16Bf4/go-evil/domains/crypto/configuration"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	notify "github.com/s9rA16Bf4/notify_handler"
)

// Generates an aes key used for encrypting/decrypting
func AES_key(value string, data_object *json.Json_t) []string {
	value = gotools.EraseDelimiter(value, []string{"\""}, -1)

	key_size := gotools.StringToInt(value)

	if key_size == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", value), "crypto.generate_key()", 1)
	}

	key := gotools.Generate_random_n_string(key_size)

	calls := configuration.Set_aes_key(key, data_object)

	return calls
}
