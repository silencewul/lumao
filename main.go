package main

import (
	"lumao/ysf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/quan1", func(c *gin.Context) {
		res := ysf.Strat(0)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": res,
		})
	})
	r.GET("/quan2", func(c *gin.Context) {
		res := ysf.Strat(1)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": res,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
