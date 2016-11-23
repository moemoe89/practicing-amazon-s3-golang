package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"practicing-amazon-s3-golang/helpers"
	"practicing-amazon-s3-golang/models"
	"practicing-amazon-s3-golang/services"
)

func Upload(c *gin.Context) {

	var uploadRequest models.UploadRequest
	c.BindJSON(&uploadRequest)

	pwd, err := os.Getwd()
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	path := pwd + "/tmp_upload"
	filename,err := helpers.WriteImage(uploadRequest.Image,path)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	err = services.AwsUpload(path,filename)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	err = helpers.DeleteFile(path+"/"+filename)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"Success upload image.",true,filename)

}

func JSONResponse(c *gin.Context,httpStatus int,message string,status bool){
	c.IndentedJSON(httpStatus, gin.H{
		"message": message,
		"status": status,
	})
}

func JSONResponseData(c *gin.Context,httpStatus int,message string,status bool,data interface{}){
	c.IndentedJSON(httpStatus, gin.H{
		"message": message,
		"status": status,
		"data": data,
	})
}