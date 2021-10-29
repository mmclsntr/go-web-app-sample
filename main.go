package main

import (
    "os"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    region := os.Getenv("REGION")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
        fmt.Println(c.Request.Header)
		c.JSON(200, gin.H{
			"region": region,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
