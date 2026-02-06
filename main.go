package main

// third party framework for REST
import (
	"github.com/gin-gonic/gin"
	"github.com/warlock1729/first-go-project/db"
	"github.com/warlock1729/first-go-project/routes"
)

func main() {
	db.InitDB()
	app := gin.Default()
	routes.RegisterRoutes(app)
	app.Run(":8080")

}
