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

type Controller struct {
	service service.StoreData
}

func NewController(svc service.StoreData) *Controller {
	return &Controller{service: svc}
}

var validate = validator.New()

func (c *Controller) GetAllData(ctx *gin.Context) {

	resp, err := c.service.GetAllData()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetDataByKey(ctx *gin.Context) {
	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	resp, err := c.service.GetDataByKey(key)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (c *Controller) SetData(ctx *gin.Context) {
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

	// TODO: is using actual db , pass the context too

	response, err := c.service.InsertData(request)
	if err != nil {
		err := fmt.Errorf("unable to insert the data")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) UpdateData(ctx *gin.Context) {
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

	response, err := c.service.UpdateDataByKey(key, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) DeleteData(ctx *gin.Context) {
	key := ctx.Param("key")

	if len(strings.TrimSpace(key)) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	err := c.service.DeleteDataByKey(key)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)

}

func (c *Controller) PushToList(ctx *gin.Context) {
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

	response, err := c.service.PushDataToList(key, request.UpdatedValue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) PopFromList(ctx *gin.Context) {
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

	response, err := c.service.PopDataFromList(key, request.UpdatedValue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
