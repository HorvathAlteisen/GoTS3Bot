package router

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type key string

const kparams key = "params"

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type Handle func(http.ResponseWriter, *http.Request, url.Values)

type Middleware func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(w, r, next)
}

func (m mware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r, m.next.ServeHTTP)
}

type Router struct {
	tree        *node
	rootHandler http.HandlerFunc
	middlewares []Middleware
}

type mware struct {
	handler Handler
	next    *mware
}

// addNode - adds a node to our tree. Will add multiple nodes if path
// can be broken up into multiple components. Those nodes will have no
// handler implemented and will fall through to the default handler.

func New(rootHandler http.HandlerFunc) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]*route)}
	return &Router{tree: &node, rootHandler: rootHandler}
}

func (r *Router) Use(m ...Middleware) {
	r.middlewares = append(r.middlewares, m...)
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc, middleware ...Middleware) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}

	r.tree.addNode(method, path, handler)
}

func (r *Router) GET(path string, handler http.HandlerFunc, middleware ...Middleware) {
	r.Handle("GET", path, handler, middleware...)
}

func (r *Router) POST(path string, handler http.HandlerFunc, middleware ...Middleware) {
	r.Handle("POST", path, handler, middleware...)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 * 1024 * 1024)

	params := req.Form

	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)

	req = contextSet(req, kparams, params) // set all the params we have collected.

	if h := node.methods[req.Method]; h != nil {
		runMiddleware(w, req, buildMList(append(r.middlewares, h.middlewares...), h.handler))
	} else {
		runMiddleware(w, req, buildMList(r.middlewares, r.rootHandler))
	}
}

func runMiddleware(w http.ResponseWriter, req *http.Request, middleware mware) {
	middleware.ServeHTTP(w, req)
}

func buildMList(middleware []Middleware, handler http.HandlerFunc) mware {
	var next mware

	if len(middleware) == 0 {
		return finalHandler(handler)
	} else if len(middleware) > 1 {
		next = buildMList(middleware[1:], handler)
	} else {
		next = finalHandler(handler)
	}

	return mware{middleware[0], &next}
}

func finalHandler(handler http.HandlerFunc) mware {
	return mware{
		Middleware(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			handler(w, r)
		}),
		&mware{},
	}
}

func contextSet(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}
	return r.WithContext(context.WithValue(r.Context(), key, val))
}
