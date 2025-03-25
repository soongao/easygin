package frame

import (
	"fmt"
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type router struct {
	handlers map[string]HandleFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			if item[0] == ':' || item[0] == '*' {
				item = string(item[0])
			}
			parts = append(parts, item)
			if item == "*" {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, fc HandleFunc) {
	parts := parsePattern(pattern)
	key := fmt.Sprintf("%v-%v", method, pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = fc
}

func parsePath(path string) []string {
	vs := strings.Split(path, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item == "*" {
				break
			}
		}
	}
	return parts
}

func (r *router) getRoute(method, path string) (string, map[string]string) {
	searchParts := parsePath(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return "", nil
	}

	n := root.search(searchParts, 0)

	if n != "" {
		parts := parsePath(n)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return "", nil
}

func (router *router) handle(ctx *Context) {
	method := ctx.Method
	path := ctx.Path
	pattern, params := router.getRoute(method, path)
	if pattern != "" {
		key := fmt.Sprintf("%v-%v", method, pattern)
		ctx.Params = params
		ctx.handles = append(ctx.handles, router.handlers[key])
	} else {
		ctx.handles = append(ctx.handles, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	ctx.Next()
}
