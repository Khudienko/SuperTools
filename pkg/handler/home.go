package handler

import (
	"github.com/Khudienko/SuperTools/pkg/views"
	"github.com/gin-gonic/gin"
)

func HomeGetHandler (c *gin.Context){
	views.GenerateHTML(c.Writer,nil,"navbar")
	views.GenerateHTML(c.Writer,nil,"home")
	views.GenerateHTML(c.Writer,nil,"footer")

}
