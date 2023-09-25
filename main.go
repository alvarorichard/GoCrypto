package main

import (
	"github.com/alvarorichard/GoCrypto/docs"
	"github.com/alvarorichard/GoCrypto/internal/utils"
	"github.com/alvarorichard/GoCrypto/internal/views"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, err := utils.NewDB()
	if err != nil {
		println("Erro ao criar o banco de dados")
		panic(err)
	}
	err = utils.MigrateDB(db)
	if err != nil {
		println("Erro ao migrar o banco de dados")
		panic(err)
	}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	views.RegisterAll(r, db)

	r.Run()
}
