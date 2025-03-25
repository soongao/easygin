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

func (rg *routerGroup) Use(middlewares ...HandleFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}
