package grappa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	ResWriter http.ResponseWriter
	Req       *http.Request
	// request
	Method string
	Path   string
	Params map[string]string
	// response
	StatusCode int
	// middleware
	handlers []HandleFunc
	index    int
	// engine pointer
	engine *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResWriter: w,
		Req:       r,
		Method:    r.Method,
		Path:      r.URL.Path,
		index:     -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) SetHeader(key string, value string) {
	c.ResWriter.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.ResWriter.WriteHeader(code)
}

func (c *Context) Param(key string) string {
	val := c.Params[key]
	return val
}

func (c *Context) PostForm(key string) string { return c.Req.FormValue(key) }

func (c *Context) Query(key string) string { return c.Req.URL.Query().Get(key) }

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.ResWriter.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.ResWriter)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

func (c *Context) Data(code int, contentType string, data []byte) {
	c.SetHeader("Content-Type", contentType)
	c.Status(code)
	c.ResWriter.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.ResWriter, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}
