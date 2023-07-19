package generate

import (
	"fmt"

	gotools "github.com/s9rA16Bf4/Go-tools"
)

func Generate_math_line() []string {
	body := []string{}
	math_operator := []string{"+", "-", "*", "/", "%"}

	op := math_operator[gotools.GenerateRandomIntBetween(0, len(math_operator))]
	a := gotools.GenerateRandomInt()
	b := gotools.GenerateRandomInt()
	c := gotools.Generate_random_string()

	body = append(body, fmt.Sprintf("%s := %d %s %d", c, a, op, b))
	body = append(body, fmt.Sprintf("payload_length += %s", c))

	return body
}
