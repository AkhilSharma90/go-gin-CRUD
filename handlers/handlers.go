package handlers

import (
	"go-gin-CRUD/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Platform *services.Services
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform *services.Services) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)
	h := Handlers{Platform: platform}

	router.GET("/", h.HelloWorld)
	return router
}

func (h *Handlers) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}
