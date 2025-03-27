package frame

type routerGroup struct {
	name   string
	engine *Engine
	isRoot bool

	middlewares []HandleFunc
}

func (rg *routerGroup) GET(pattern string, fc HandleFunc) {
	rg.engine.router.addRoute("GET", rg.name+pattern, fc)
}

func (rg *routerGroup) POST(pattern string, fc HandleFunc) {
	rg.engine.router.addRoute("POST", rg.name+pattern, fc)
}

func (rg *routerGroup) PUT(pattern string, fc HandleFunc) {
	rg.engine.router.addRoute("PUT", rg.name+pattern, fc)
}

func (rg *routerGroup) DELETE(pattern string, fc HandleFunc) {
	rg.engine.router.addRoute("DELETE", rg.name+pattern, fc)
}

func (rg *routerGroup) Use(middlewares ...HandleFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

type RESTful interface {
	Create(*Context)
	Query(*Context)
	Update(*Context)
	Delete(*Context)
}

func (rg *routerGroup) REST(pattern string, rt RESTful) {
	rg.GET(pattern, rt.Query)
	rg.POST(pattern, rt.Create)
	rg.PUT(pattern, rt.Update)
	rg.DELETE(pattern, rt.Delete)
}

type BaseREST struct{}

func (base *BaseREST) Create(c *Context) { c.JSON(200, "RESTful POST") }

func (base *BaseREST) Query(c *Context) { c.JSON(200, "RESTful GET") }

func (base *BaseREST) Update(c *Context) { c.JSON(200, "RESTful PUT") }

func (base *BaseREST) Delete(c *Context) { c.JSON(200, "RESTful DELETE") }
