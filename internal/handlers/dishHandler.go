package handlers

import (
	"SaltAIdDishes/internal/databaseModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Generate(c *gin.Context) {
	type myJSON struct {
		Params []string `json:"params"`
	}
	var jsonInput myJSON
	if err := c.BindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result []string
	for _, str := range jsonInput.Params {
		if str != "" {
			result = append(result, str)
		}
	}
	jsonInput.Params = result
	found, err := databaseModels.Dishes.GetWithParams(jsonInput.Params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"english":     found.Name,
			"description": found.Description,
			"ingredients": found.Ingredients,
			"recipe":      found.Recipe,
			"image":       found.Url,
			"video":       found.Link,
			"russian":     found.RussianName,
		})
	}
}
