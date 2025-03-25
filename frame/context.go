package frame

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Params map[string]string
	Path   string
	Method string

	StatusCode int

	handles []HandleFunc
	index   int
}

type H map[string]any

func newContext(w http.ResponseWriter, q *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    q,
		Path:   q.URL.Path,
		Method: q.Method,

		handles: make([]HandleFunc, 0),
		index:   -1,
	}
}

func (ctx *Context) Param(key string) string {
	return ctx.Params[key]
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) JSON(code int, val any) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)

	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(val); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) Next() {
	ctx.index++
	for ; ctx.index < len(ctx.handles); ctx.index++ {
		ctx.handles[ctx.index](ctx)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handles)
	c.JSON(code, H{"message": err})
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
