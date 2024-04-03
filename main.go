package main

import (
	"gooob/gooob"
	"net/http"
)

func main() {
	r := gooob.Default()
	r.GET("/", func(c *gooob.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gooob.Context) {
		names := []string{"gooob"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
