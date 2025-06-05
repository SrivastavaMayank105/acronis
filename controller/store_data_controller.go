package controller

import (
	"acronis/mapper"
	"acronis/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetAllData(ctx *gin.Context) {
	svc := service.NewInMemoryStore()
	resp, err := svc.GetAllData()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func GetDataByKey(ctx *gin.Context) {
	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	svc := service.NewInMemoryStore()
	resp, err := svc.GetDataByKey(key)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, resp)

}

func SetData(ctx *gin.Context) {
	var request mapper.DataInfo

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := validate.Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// TODO: check if the data is empty or not , if so return error

	svc := service.NewInMemoryStore()
	response, err := svc.InsertData(request)
	if err != nil {
		err := fmt.Errorf("unable to insert the data")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func UpdateData(ctx *gin.Context) {
	var request mapper.DataInfo

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := validate.Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	svc := service.NewInMemoryStore()
	response, err := svc.UpdateDataByKey(key, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func DeleteData(ctx *gin.Context) {
	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	svc := service.NewInMemoryStore()
	err := svc.DeleteDataByKey(key)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)

}

func PushToList(ctx *gin.Context) {
	var request mapper.UpdateListData

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := validate.Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	svc := service.NewInMemoryStore()
	response, err := svc.PushDataToList(key, request.UpdatedValue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func PopFromList(ctx *gin.Context) {
	var request mapper.UpdateListData

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := validate.Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	svc := service.NewInMemoryStore()
	response, err := svc.PopDataFromList(key, request.UpdatedValue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
