package controllers

import (
	"net/http"

	"bitbucket.org/joisandresky/qurban-v2/models"
	"github.com/gin-gonic/gin"
)

// GetMitra - route handler for get all mitra
func GetMitra(c *gin.Context) {
	var MitraModel models.Mitra
	mitras, err := MitraModel.GetMitra()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "an Error Occured trying to get Mitras",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mitras": mitras,
	})
}
