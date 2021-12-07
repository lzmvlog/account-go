package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}
