package handlers

import (
	"net/http"

	"go-gin-CRUD/domain"
	"go-gin-CRUD/services"

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
	router := gin.Default()

	h := Handlers{Platform: platform}

	book := router.Group("/book")
	book.GET("/", h.GetAllBooks)
	book.GET("/:id", h.GetBook)
	book.POST("/", h.CreateBook)
	book.PUT("/:id", h.UpdateBook)
	book.DELETE("/:id", h.DeleteBook)

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

func (h *Handlers) CreateBook(ctx *gin.Context) {
	var response Response
	var request domain.Book
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book, code, err := h.Platform.CreateBook(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) UpdateBook(ctx *gin.Context) {
	var response Response
	var request domain.Book
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	book, err := h.Platform.UpdateBook(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) DeleteBook(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.DeleteBook(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
