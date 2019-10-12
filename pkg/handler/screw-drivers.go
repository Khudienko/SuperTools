package handler

import (
	"github.com/Khudienko/SuperTools/pkg/views"
	"github.com/gin-gonic/gin"
)

func ScrewDriversGetHandler(c *gin.Context){
	views.GenerateHTML(c.Writer,nil,"navbar")
	views.GenerateHTML(c.Writer,nil,"screw_drivers")
	views.GenerateHTML(c.Writer,nil,"footer")

}
