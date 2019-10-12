package handler

import (
	"github.com/Khudienko/SuperTools/pkg/models"
	"github.com/Khudienko/SuperTools/pkg/views"
	"github.com/gin-gonic/gin"
)


type ProdPageView struct {
	Product models.Product
}

func (m *ModelsMethods) ProductGetHandler(c *gin.Context){

	//name:=m.Product.GetProdName(1)
	prod:=ProdPageView{
		Product:m.Product.GetProduct(1),
	}

	views.GenerateHTML(c.Writer,nil,"navbar")
	views.GenerateHTML(c.Writer,prod,"products")
	views.GenerateHTML(c.Writer,nil,"footer")

}
