package generate

import tools "github.com/s9rA16Bf4/Go-tools"

func Generate_function_body(return_type string) []string {
	body := []string{}

	body_lines := tools.GenerateRandomIntBetween(1, 32)
	body = append(body, "payload_body := \"A\"")
	body = append(body, "payload_length := 0")

	for line := 0; line < body_lines; line++ {
		line_type := tools.GenerateRandomIntBetween(0, 3)

		switch line_type {
		case 1: // String
			body = append(body, Generate_string_line()...)

		case 2: // Math
			body = append(body, Generate_math_line()...)
		}

	}
	body = append(body, "fmt.Sprintf(\"%s\", payload_body)", "fmt.Sprintf(\"%d\", payload_length)")

	// Generate return value/values
	body = append(body, Generate_return_string_line(return_type))

	return body
}
