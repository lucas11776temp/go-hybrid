package courses

import (
	"net/http"
	validator "test/src/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// c.JSON(http.StatusOK, [])
}

func Store(c *gin.Context) {
	validator := validator.Validator{
		Rules: validator.Rules{
			"title":       {"required", "min:3", "max:20"},
			"description": {"required", "min:3", "max:50"},
		},
	}

	ok := validator.Validate(c.Request)

	if !ok {
		c.JSON(http.StatusUnprocessableEntity, validator.Errors())
		return
	}

	// c.JSON(http.StatusOK, [])
}

func View(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Destroy(c *gin.Context) {

}
