package grappa

import (
	"net/http"
	"strings"
)

type router struct {
	roots   map[string]*node
	handles map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*node), // a method correspond to a trie
		handles: make(map[string]HandleFunc),
	}
}

func dividePath(path string) []string {
	raw := strings.Split(path, "/")

	parts := make([]string, 0)
	for _, part := range raw {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) routeRegister(method string, path string, handler HandleFunc) {
	parts := dividePath(path)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(path, parts, 0)

	key := method + "-" + path
	r.handles[key] = handler
}

func (r *router) parseRoute(method string, path string) (*node, map[string]string) {
	rawParts := dividePath(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	end := root.search(rawParts, 0)

	if end != nil {
		parts := dividePath(end.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = rawParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(rawParts[index:], "/")
				break
			}
		}

		return end, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	end, params := r.parseRoute(c.Method, c.Path)

	if end != nil {
		c.Params = params
		key := c.Method + "-" + end.path
		c.handlers = append(c.handlers, r.handles[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
		})
	}
	c.Next()

}

func (r *router) checkPaths(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	paths := make([]*node, 0)
	root.traverse(&paths)
	return paths
}
