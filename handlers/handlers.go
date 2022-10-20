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
	book := router.Group("/book")
	book.GET("/", h.GetAllBooks)
	book.GET("/:id", h.GetBook)
	router.GET("/", h.HelloWorld)
	return router
}

func (h *Handlers) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func (h *Handlers) GetAllBooks(ctx *gin.Context) {
	var response Response

	books, err := h.Platform.GetAllBooks()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = books
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) GetBook(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	book, err := h.Platform.GetBook(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
