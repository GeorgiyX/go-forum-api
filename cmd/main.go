package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const VolFile = "/vol/file.txt"
const NotVolFile = "/tmp/file.txt"

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.GET("/save/vol", func(c *gin.Context) {
		writeHandler(c, VolFile)
	})
	r.GET("/save/notvol", func(c *gin.Context) {
		writeHandler(c, NotVolFile)
	})
	r.GET("/read/vol", func(c *gin.Context) {
		readHandler(c, VolFile)
	})
	r.GET("/read/notvol", func(c *gin.Context) {
		readHandler(c, NotVolFile)
	})
	err := r.Run()
	if err != nil {
		log.Fatal("Gin not started!")
		return
	}
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func writeHandler(c *gin.Context, filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
	}()

	if _, err = file.Write([]byte(fmt.Sprintf("hello volume: %v\n", time.Now()))); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "write ok",
	})
}

func readHandler(c *gin.Context, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(bytes),
	})
}
