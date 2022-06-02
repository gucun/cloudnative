package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// router := gin.Default()
	router := gin.Default()
	router.GET("healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "200", c.Request.Header)

	})
	router.Run()

}
