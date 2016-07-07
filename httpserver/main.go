package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func postMan(c *gin.Context) {

	in := struct {
		Cert string `json:"cert"`
		Key  string `json:"key"`
	}{}

	in1, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(in1, &in)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(in.Cert)

	c.JSON(200, gin.H{
		"msg": "ok",
	})

}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/postMan", postMan)
	r.Run()
	fmt.Println("help")

}
