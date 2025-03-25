package frame

import (
	"log"
	"net/http"
	"strings"
)

type Engine struct {
	*routerGroup
	router *router
	groups []*routerGroup
}

func NewEngine() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.routerGroup = &routerGroup{
		engine: engine,
		isRoot: true,
	}
	engine.groups = append(engine.groups, engine.routerGroup)
	return engine
}

func Default() *Engine {
	engine := NewEngine()
	engine.Use(Logger(), Recover())
	return engine
}

func (e *Engine) Group(prefix string) *routerGroup {
	group := &routerGroup{
		name:   e.name + prefix,
		engine: e,
	}
	e.groups = append(e.groups, group)
	return group
}

func (e *Engine) Run(addr string) {
	log.Printf("Listen and Server address: %s", addr)
	log.Fatalln(http.ListenAndServe(addr, e))
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middlewares := make([]HandleFunc, 0)
	for _, g := range e.groups {
		if strings.HasPrefix(r.URL.Path, g.name) {
			middlewares = append(middlewares, g.middlewares...)
		}
	}
	ctx := newContext(w, r)
	ctx.handles = middlewares
	e.router.handle(ctx)
}
