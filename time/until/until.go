package until

import (
	"fmt"
	"regexp"
	"time"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	notify "github.com/DennisTheodoreNedry/notify_handler"
)

const (
	GRAB_FULL_DATE = "([0-9]{0,4})/([0-9]{0,2})/([0-9]{0,2})-([0-9]{0,2}):([0-9]{0,2})" // YYYY/MM/DD-hh:mm
	GRAB_HOUR_MIN  = "([0-9]{0,2}):([0-9]{0,2})"                                        // hh:mm
)

// Makes the malware wait until the yyyy-mm-dd-hh-mm has been reached
func Until(value string, data_object *json.Json_t) []string {
	function_call := "Until"

	regex := regexp.MustCompile(GRAB_FULL_DATE)
	result_full := regex.FindAllStringSubmatch(value, -1)

	regex2 := regexp.MustCompile(GRAB_HOUR_MIN)
	result_hour_min := regex2.FindAllStringSubmatch(value, -1)

	if len(result_full) == 0 && len(result_hour_min) == 0 {
		notify.Error("Failed to find a valid time format!", "time.Until", 1)
	}

	year, month, day := time.Now().Date()

	parameters := []string{
		gotools.IntToString(year),
		gotools.IntToString(int(month)),
		gotools.IntToString(day),
	}

	if len(result_full) > 0 {
		parameters = []string{result_full[0][1], result_full[0][2], result_full[0][3], result_full[0][4], result_full[0][5]}
	} else {
		parameters = append(parameters, result_hour_min[0][1], result_hour_min[0][2]) // Grab the hour and minute
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 []int", "repr_3 []int", "repr_4 []int", "repr_5 []int"},

		Gut: []string{
			"year := gotools.StringToInt(spine.variable.get(spine.alpha.construct_string(repr_1)))",
			"month := gotools.StringToInt(spine.variable.get(spine.alpha.construct_string(repr_2)))",
			"day := gotools.StringToInt(spine.variable.get(spine.alpha.construct_string(repr_3)))",
			"hour := gotools.StringToInt(spine.variable.get(spine.alpha.construct_string(repr_4)))",
			"minute := gotools.StringToInt(spine.variable.get(spine.alpha.construct_string(repr_5)))",
			"c_now := time.Now()",
			"for year != c_now.Year() || month != int(c_now.Month()) || day != c_now.Day() || hour != c_now.Hour() || minute != c_now.Minute() {",
			"time.Sleep(5 * (10 ^ 9))",
			"c_now = time.Now()",
			"}",
		}})

	data_object.Add_go_import("time")

	constructed_parameters := []string{
		data_object.Generate_int_array_parameter(parameters[0]),
		data_object.Generate_int_array_parameter(parameters[1]),
		data_object.Generate_int_array_parameter(parameters[2]),
		data_object.Generate_int_array_parameter(parameters[3]),
		data_object.Generate_int_array_parameter(parameters[4]),
	}

	return []string{fmt.Sprintf("%s(%s, %s, %s, %s, %s)", function_call, constructed_parameters[0], constructed_parameters[1], constructed_parameters[2],
		constructed_parameters[3], constructed_parameters[4])}
}
