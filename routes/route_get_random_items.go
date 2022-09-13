package routes

import (
	"net/http"

	"github.com/allefts/minedle/controllers"
	"github.com/gin-gonic/gin"
)

func GetRandomItemRoute(c *gin.Context) {

	Items, err := controllers.GetRandomItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, Items)
		return
	}
}
