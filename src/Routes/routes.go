package Routes

import (
	"github.com/bezaeel/rest-api-mysql-gin/src/Controllers"
	"github.com/gin-gonic/gin"
)

//configure routes
func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/api")
	{
		grpl.GET("contacts/all", Controllers.GetAllContacts)
		grpl.POST("contacts/add", Controllers.AddContact)
		grpl.GET("GetContact/:id", Controllers.GetContactById)
		grpl.PUT("contacts/:id/edit", Controllers.UpdateContact)
		grpl.DELETE("contacts/:id/remove", Controllers.DeleteContact)
	}
	return r
}
