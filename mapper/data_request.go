package mapper

import "github.com/go-playground/validator/v10"

type DataInfo struct {
	Data interface{} `json:"data" validate:"required"`
}

var validate = validator.New()

func ValidateDataInfo(dataInfo *DataInfo) error {
	return validate.Struct(dataInfo)
}

type UpdateListData struct {
	UpdatedValue interface{} `json:"uptvalue"`
}

func ValidateUpdateListData(uptLstData *UpdateListData) error {
	return validate.Struct(uptLstData)
}
