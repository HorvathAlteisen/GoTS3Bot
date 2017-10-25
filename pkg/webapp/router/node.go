package router

import (
	"net/http"
	"net/url"
	"strings"
)

type route struct {
	handler     http.HandlerFunc
	middlewares []Middleware
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]*route
}

func (n *node) addNode(method, path string, handler http.HandlerFunc, middleware ...Middleware) {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	for {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 { // update an existing node.
			r := route{handler: handler}
			r.middlewares = append(r.middlewares, middleware...)
			aNode.methods[method] = &r
			return
		}
		newNode := node{component: component, isNamedParam: false, methods: make(map[string]*route)}

		if len(component) > 0 && component[0] == ':' { // check if it is a named param.
			newNode.isNamedParam = true
		}
		if count == 1 { // this is the last component of the url resource, so it gets the handler.
			r := route{handler: handler}
			r.middlewares = append(r.middlewares, middleware...)
			newNode.methods[method] = &r
		}
		aNode.children = append(aNode.children, &newNode)
		count--
		if count == 0 {
			break
		}
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 { // no children, then bail out.
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 { // http://xkcd.com/1270/
					return child.traverse(next, params) // tail recursion is it's own reward.
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
