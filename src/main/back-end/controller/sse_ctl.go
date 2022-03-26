package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

/**
    controller
    @author: roccoshi
    @desc: For Sever Sent Events (SSE)
**/

func SSETestAction(c *gin.Context) {
	streamChan := make(chan int)
	go func () {
		defer close(streamChan)
		for i := 0; i < 10; i ++ {
			streamChan <- i
			time.Sleep(3 * time.Second)
		}
	}()
	c.Stream(func (w io.Writer) bool {
		if msg, ok := <-streamChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}