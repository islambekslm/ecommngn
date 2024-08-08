package handlers

import (
	"ecommngn/internal/infrastructure/inmemdb"
	"ecommngn/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	_repo    = inmemdb.NewDB[product.Product]()
	_service = product.NewManager(_repo)
)

func RegisterOrderRoutes(router *gin.Engine) {
	r := router.Group("/order")
	r.GET("/:id", GetOrder)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	res, err := _service.Get(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, res)
}
