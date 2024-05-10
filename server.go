package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/example", func(c *gin.Context) {
		var realData interface{}
		r, err := Select()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if r == nil {
			realData = make([]Data, 0)
		} else {
			realData = r
		}

		c.JSON(200, Response{
			Data: realData,
			Meta: "Success",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

type Data struct {
	Field1 string
}

func Select() (*Data, error) {
	return nil, nil
}

type Response struct {
	Data interface{}
	Meta string
}
