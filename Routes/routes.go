package Routes

import (
	"github.com/bezaeel/rest-api-mysql-gin/Controllers"
	"github.com/gin-gonic/gin"
)

//configure routes
func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/contacts")
	{
		grpl.GET("all", Controllers.GetAllContacts)
		grpl.POST("add", Controllers.AddContact)
		grpl.PUT("edit/:id", Controllers.UpdateContact)
		grpl.DELETE("remove/:id", Controllers.DeleteContact)
	}
	return r
}
