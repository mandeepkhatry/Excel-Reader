package validation

//Validate config according to the standard protocol

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Sheets struct {
	Sheet_Name    string                 `json:"sheet_name" validate:"required,alphanum,min=1"`
	Sheet_Type    string                 `json:"sheet_type" validate:"required,alpha,eq=Tabular|eq=Columnar|eq=ColTab|eq=tabular|eq=columnar|eq=coltab"`
	Offset_Status bool                   `json:"offset_status" validate:"required"`
	Offset_Value  int                    `json:"offset_value" validate:"numeric,gte=0"`
	Columns       map[string]interface{} `json:"columns" validate:"required,dive,keys,alpha,len=1,endkeys"`
	Rows          map[string]interface{} `json:"rows" validate:"required,dive,keys,numeric,endkeys"`
	Rules         string                 `json:"rules" validate:"required"`
}

type Config struct {
	Config string    `json:"config" validate:"required,alphanum,min=2"`
	Sheets []*Sheets `json:"sheets" validate:"required,dive"`
}

func ValidateFileConfig(config map[string]interface{}) bool {
	var configuration Config
	data, _ := json.Marshal(config)
	json.Unmarshal(data, &configuration)
	for _, v := range configuration.Sheets {
		fmt.Println(v)
	}
	v := validator.New()
	err := v.Struct(configuration)

	if err != nil {
		return false
	}
	return true
}
