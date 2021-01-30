package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi ROtHi..",
		})
	})
	router.GET("/me/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"Welcome": name,
		})
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.POST("/form_post", func(c *gin.Context) {
		type LoginCommand struct {
			Username string 
			Password string 
		}
		var loginCmd LoginCommand
		c.BindJSON(&loginCmd)
		c.JSON(http.StatusOK, gin.H{"user": loginCmd})
	})

	router.Run() 
	//without gin reload (go run main.go)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//with gin realod (gin run main.go)
	//localhost:3000 (default port)
}