package routes

import (
	"net/http"

	"github.com/allefts/minedle/controllers"
	"github.com/gin-gonic/gin"
)

type Response struct {
	ResultItem      interface{}
	IngredientItems []controllers.ItemsUsed
}

// Response format is very complex, maybe need to reformat or just handle it on the frontend
func GetItemRoute(c *gin.Context) {
	Item, err := controllers.HandleItem()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	Items, err := controllers.GetRandomItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	res := Response{ResultItem: Item, IngredientItems: Items}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}
