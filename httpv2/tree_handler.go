package httpv2

import (
	"net/http"
	"strings"
)

type HandlerBasedOnTree struct {
	root *node
}
type node struct {
	path     string
	children []*node
	//如果这是叶子节点 那么匹配上之后就可以调用该方法
	handler HandleFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	url := strings.Trim(c.R.URL.Path, "/")
	paths := strings.Split(url, "/")
	cur := h.root
	for _, path := range paths {
		matchChidl, found := h.findMatchChild(cur, path)
		if !found {
			c.W.WriteHeader(http.StatusNotFound)
			c.W.Write([]byte("404 Not Found!!!"))
			return
		}
		cur = matchChidl
	}
	if cur.handler == nil {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("404 Not FFFFFound!!!"))
		return
	}
	cur.handler(c)
}

func (h *HandlerBasedOnTree) Route(method string, pat string, hand HandleFunc) {
	pat = strings.Trim(pat, "/")
	paths := strings.Split(pat, "/")
	cur := h.root
	for i, path := range paths {
		mathChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = mathChild
		} else {
			h.createSubTree(cur, paths[i:], hand)
			return
		}
	}
}
func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handleFunc HandleFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handleFunc
}
func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 8),
	}
}
func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range root.children {
		if child.path == path && child.path != "*" {
			return child, true
		}
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}
