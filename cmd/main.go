package main

import (
	"github.com/Khudienko/SuperTools/gomigrations"
	"github.com/Khudienko/SuperTools/pkg/driver"
	"github.com/Khudienko/SuperTools/pkg/handler"
	"github.com/Khudienko/SuperTools/pkg/logger"
	"github.com/Khudienko/SuperTools/pkg/models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func main() {
	db := driver.ConnectDB()
	err := gomigrations.Migrate(db)
	if err != nil {
		logger.FatalError(err, "Migration failed.\n")
	}


	connection:=models.Connection{DB:db}
	controll:=handler.ModelsMethods{
		Product:&connection,
	}





	r := gin.Default()
	r.GET("/",handler.HomeGetHandler)
	r.GET("/powertools",handler.PowerToolGetHandler)
	r.GET("/powertools/sw-drivers",handler.ScrewDriversGetHandler)
	r.GET("/product",controll.ProductGetHandler)


	r.Use(static.Serve("/", static.LocalFile("./web/templates/", true)))
	r.Run(":8080")
}
