package main

import (
	"github.com/bitmon-world/bitmon-api/controllers"
	"github.com/bitmon-world/bitmon-api/models"
	"github.com/bitmon-world/bitmon-api/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	App := GetApp()
	err := App.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// GetApp is used to wrap all the additions to the GIN API.
func GetApp() *gin.Engine {
	App := gin.Default()
	App.Use(cors.Default())
	ApplyRoutes(App)
	return App
}

func ApplyRoutes(r *gin.Engine) {
	r.Static("/img", "./img")
	ctrl := controllers.BitmonController{
		Elements: models.Elements,
		Bitmons:  models.Bitmons,
	}
	api := r.Group("/")
	{
		api.GET("/mon/single/:id", func(c *gin.Context) { callWrapper(c, ctrl.GetMon) })
		api.GET("/elements/single/:id", func(c *gin.Context) { callWrapper(c, ctrl.GetElement) })
	}
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

func callWrapper(c *gin.Context, method func(params types.ReqParams) (interface{}, error)) {
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	params := types.ReqParams{
		ID:   id,
		Body: body,
	}
	res, err := method(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, res)
		return
	}
}
