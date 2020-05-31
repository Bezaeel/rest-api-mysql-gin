package Controllers

import (
	"fmt"
	"net/http"

	"github.com/bezaeel/rest-api-mysql-gin/src/Models"
	"github.com/bezaeel/rest-api-mysql-gin/src/Services"
	"github.com/bezaeel/rest-api-mysql-gin/src/Communication"
	"github.com/gin-gonic/gin"
)

//var apiResponse = new(ApiResponse.Response)

//GetAllContacts
func GetAllContacts(c *gin.Context) {
	var contact []Models.Contact
	// var apiResponse Communication.apiResponse.apiResponse
	var apiResponse = new(ApiResponse.Response)

	err := Services.GetAllContacts(&contact)
	if err != nil {

		c.JSON(http.StatusNotFound, apiResponse)
		return
	} else {
		apiResponse.IsSuccess = true
		apiResponse.Message = "Success"
		apiResponse.Data = contact
		c.JSON(http.StatusOK, apiResponse)
		return
	}
}

func GetContactById(c *gin.Context) {
	var contact Models.Contact
	var apiResponse = new(ApiResponse.Response)
	id := c.Params.ByName("id")
	err := Services.GetContactById(id, &contact)
	if err != nil {
		// apiResponse.IsSuccess = false
		apiResponse.Message = err.Error()
		c.JSON(http.StatusNotFound, apiResponse)
		return
	} else {
		apiResponse.IsSuccess = true
		apiResponse.Message = "Success"
		apiResponse.Data = contact
		c.JSON(http.StatusOK, apiResponse)
		return
	}
}

func AddContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	var apiResponse = new(ApiResponse.Response)
	err := Services.CreateContact(&contact)
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
		apiResponse.Message = err.Error()
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	} else {
		apiResponse.IsSuccess = true
		apiResponse.Message = "Success"
		apiResponse.Data = contact
		c.JSON(http.StatusOK, apiResponse)
		return
	}
}

func UpdateContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	var apiResponse = new(ApiResponse.Response)
	err := Services.GetContactById(id, &contact)
	if err != nil {
		apiResponse.Message = err.Error()
		c.JSON(http.StatusNotFound, apiResponse)
		return
	} else {
		c.BindJSON(&contact)
		err := Services.UpdateContact(id, &contact)
		if err != nil {
			apiResponse.Message = err.Error()
			c.JSON(http.StatusInternalServerError, apiResponse)
			return
		}
		apiResponse.IsSuccess = true
		apiResponse.Message = "contact with id: " + id + " updated successfully"
		apiResponse.Data = contact
		c.JSON(http.StatusOK, apiResponse)
		return
	}
}

func DeleteContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	var apiResponse = new(ApiResponse.Response)
	err := Services.GetContactById(id, &contact)
	if err != nil {
		apiResponse.Message = err.Error()
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}
	err = Services.DeleteContact(id, &contact)
	if err != nil {
		apiResponse.Message = err.Error()
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}
	apiResponse.IsSuccess = true
	apiResponse.Message = "contact with id: " + id + " deleted successfully"
	apiResponse.Data = contact
	c.JSON(http.StatusOK, apiResponse)
	return
}
