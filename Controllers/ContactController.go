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
		return
	} else {
		c.JSON(http.StatusOK, contact)
		return
	}
}

func AddContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	err := Services.CreateContact(&contact)
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.JSON(http.StatusOK, contact)
		return
	}
}

func UpdateContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Services.GetContactById(id, &contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.BindJSON(&contact)
		err := Services.UpdateContact(id, &contact)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, contact)
		return
	}
}

func DeleteContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Services.GetContactById(id, &contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = Services.DeleteContact(id, &contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	message := "contact with id: " + id + " deleted successfully"
	c.JSON(http.StatusOK, gin.H{"message": message})
	return
}
