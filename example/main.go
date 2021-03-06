package main

import (
	"bytes"

	"github.com/vicanso/elton"

	jp "github.com/vicanso/elton-json-picker"
)

func main() {

	e := elton.New()

	e.Use(jp.NewDefault("_fields"))

	// http://127.0.0.1:7001/?_fields=foo,id
	e.GET("/", func(c *elton.Context) (err error) {
		c.SetHeader(elton.HeaderContentType, elton.MIMEApplicationJSON)
		c.BodyBuffer = bytes.NewBufferString(`{
			"foo": "bar",
			"id": 1,
			"price": 1.21
		}`)
		return
	})

	err := e.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}
