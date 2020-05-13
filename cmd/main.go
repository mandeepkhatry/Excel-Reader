package main

import (
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

				"rules": "!isNull(A) && !isNull(B)",
			},
		},
	}

	fmt.Println("Check validation : ", validation.ValidateFileConfig(config))

}
