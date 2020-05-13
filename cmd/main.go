package main

import (
	"Excel-Reader/utility"
	"Excel-Reader/validation"
	"fmt"
)

func main() {

	var config = map[string]interface{}{
		"config": "test",
		"sheets": []interface{}{
			map[string]interface{}{
				"sheet_name":    "hello",
				"sheet_type":    "columnar",
				"offset_status": true,
				"offset_value":  0,
				"columns": map[string]interface{}{
					"A": "name",
					"B": "age",
				},
				"rows": map[string]interface{}{
					"1": "name",
					"2": "age",
				},

				"rules": "!IsNull(A) && IsNull(B) && B>10 && IsNull(12)",
			},
		},
	}

	result, err := utility.SeperateExpressions("!IsNull(A) && IsNull(B) && B>10 && IsNull(12)")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	fmt.Println("Check validation : ", validation.ValidateConfig(config))

}
