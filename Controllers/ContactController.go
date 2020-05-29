package Controllers

import (
	"fmt"
	"net/http"
	"github.com/bezaeel/rest-api-mysql-gin/Models"
	"github.com/bezaeel/rest-api-mysql-gin/Services"
	"github.com/gin-gonic/gin"
)

//GetAllContacts
func GetAllContacts(c *gin.Context) {
	var contact []Models.Contact
	err := Services.GetAllContacts(&contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}

func AddContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	err := Services.CreateContact(&contact)
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}
