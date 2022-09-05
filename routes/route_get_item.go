package routes

import (
	"net/http"

	"github.com/allefts/minedle/controllers"
	"github.com/gin-gonic/gin"
)

// Response format is very complex, maybe need to reformat or just handle it on the frontend
func GetItemRoute(c *gin.Context) {
	Item, err := controllers.HandleItem()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, Item)
		return
	}
}
