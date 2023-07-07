package generate

import (
	"fmt"

	tools "github.com/s9rA16Bf4/Go-tools"
)

func Generate_math_line() []string {
	body := []string{}
	math_operator := []string{"+", "-", "*", "/", "%"}

	op := math_operator[tools.GenerateRandomIntBetween(0, len(math_operator))]
	a := tools.GenerateRandomInt()
	b := tools.GenerateRandomInt()
	c := tools.Generate_random_string()

	body = append(body, fmt.Sprintf("%s := %d %s %d", c, a, op, b))
	body = append(body, fmt.Sprintf("payload_length += %s", c))

	return body
}
