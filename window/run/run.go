package run

import (
	"fmt"
	"os"
	"strings"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Drops all the needed code from the json strucutre into one function
func Run(data_object *json.Json_t) []string {
	function_call := "Run()"
	final_content := ""

	if data_object.Index_file == "" {
		html_content := strings.Join(data_object.Html_gut, "\n")
		js_content := strings.Join(data_object.Js_gut, "\n")
		css_content := strings.Join(data_object.Css_gut, "\n")

		final_content = fmt.Sprintf(`
		<head>
			<title>%s</title>
			<style>%s</style>
			<script>%s</script>
		</head>
		<body>
		%s
		</body>
		`, data_object.Title, css_content, js_content, html_content)
	} else {
		gut, err := os.ReadFile((data_object.Index_file))

		if err != nil {
			notify.Error(err.Error(), "window.Run()", 1)
		}

		final_content = string(gut)
	}

	binding := ""
	for key := range data_object.Bind_gut {
		binding += fmt.Sprintf("w.Bind(%s, %s)\n", key, data_object.Bind_gut[key])
	}

	body := []string{
		fmt.Sprintf("win, err := lorca.New(fmt.Sprintf(\"data:text/html,%%s\", url.PathEscape(`%s`)), \"\", %d, %d)", final_content, data_object.Width, data_object.Height),
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
	}

	for key := range data_object.Bind_gut {
		body = append(body, fmt.Sprintf("win.Bind(%s, %s)\n", key, data_object.Bind_gut[key]))
	}

	body = append(body, "defer win.Close()", "<-win.Done()")

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "", Parameters: []string{}, Gut: body})

	data_object.Add_go_import("github.com/DennisTheodoreNedry/lorca")
	data_object.Add_go_import("net/url")
	data_object.Add_go_import("github.com/DennisTheodoreNedry/notify_handler")
	data_object.Add_go_import("fmt")

	return []string{function_call}
}
