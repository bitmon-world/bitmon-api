package main

import (
	"github.com/bitmon-world/bitmon-api/controllers"
	"github.com/bitmon-world/bitmon-api/types"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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
	api := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("AUTH_USERNAME"): os.Getenv("AUTH_PASSWORD"),
	}))
	{
		store := persistence.NewInMemoryStore(time.Hour)
		ctrl := controllers.NewBitmonController(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_NAME"))

		// General Information
		api.GET("/mon/single/:id", cache.CachePage(store, time.Minute*10, func(c *gin.Context) { callWrapper(c, ctrl.GetMonInfo) }))
		api.GET("/mon/list", cache.CachePage(store, time.Minute*10, func(c *gin.Context) { callWrapper(c, ctrl.GetMonList) }))
		api.POST("/mon/add", func(c *gin.Context) { callWrapper(c, ctrl.GetMonList) })
		api.GET("/elements/list", cache.CachePage(store, time.Minute*10, func(c *gin.Context) { callWrapper(c, ctrl.GetElementsList) }))
		api.GET("/elements/single/:id", cache.CachePage(store, time.Minute*10, func(c *gin.Context) { callWrapper(c, ctrl.GetMonList) }))
		api.GET("/elements/image/:id", cache.CachePage(store, time.Minute*10, func(c *gin.Context) { callWrapper(c, ctrl.GetMonList) }))
		api.POST("/elements/add", func(c *gin.Context) { callWrapper(c, ctrl.AddElement) })

		// Adventure algorithm
		api.POST("/adventure", func(c *gin.Context) { callWrapper(c, ctrl.CalcAdventure) })
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
