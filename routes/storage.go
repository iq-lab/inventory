package routes

import (
	"inventory/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetParentStorages(c *gin.Context) {
	data, err := db.GetParentStorages()
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, data)
}
