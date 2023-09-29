package main

import (
	"github.com/alvarorichard/GoCrypto/internal/utils"
	"github.com/alvarorichard/GoCrypto/internal/views"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	c:=cors.DefaultConfig()
	c.AllowOrigins = []string{"*"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(c))
	//r.Use(cors.Default())

	views.RegisterAll(r, db)

	r.Run()
}
