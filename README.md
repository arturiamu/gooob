## gooob

gooob is a simple HTTP web framework written in Go (Golang). Used to learn and understand the principles of web frameworks, the code comes from [geektutu/7days-golang](https://github.com/geektutu/7days-golang)

## Features

- Dynamic routing
- Middleware support
- Crash-free
- Routes grouping
- Static files serving

## Usage

```go
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
```
