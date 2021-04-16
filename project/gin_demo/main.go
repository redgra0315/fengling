package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func func1(c *gin.Context) {
	fmt.Println("func1  before")
	c.Next()
	fmt.Println("func1 after")
}

func func2(c *gin.Context) {
	fmt.Println("func2  before")
	c.Next()
	fmt.Println("func2 after")
}

func func3(c *gin.Context) {
	fmt.Println("func3  before")
	c.Next()
	fmt.Println("func3 after")
}

func func4(c *gin.Context) {
	fmt.Println("func4  before")
	//c.Next()
	c.Set("name", "hangzhou")
	fmt.Println("func4 after")
}

func func5(c *gin.Context) {
	fmt.Println("func5  before")
	value, ok := c.Get("name")
	if ok {
		vStr := value.(string) // 类型转换
		fmt.Println(vStr)
	}
	fmt.Println("func5 after")
}

func main() {
	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	shopGroup := r.Group("/shop", func1, func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index", func4, func5)
	}
	r.Run()
}
