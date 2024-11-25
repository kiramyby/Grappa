package grappa

import "strings"

type node struct {
	path     string
	part     string
	children []*node
	isFuzz   bool
}

// matchChild return the *node value of the matching child.
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isFuzz {
			return child
		}
	}
	return nil
}

// matchChildren return all the *node values of the matching child.
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isFuzz {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// insert a node chain into the tree. into insert means the node exist or has been insert
// path is the complete path, and parts are the divided path you want to insert
// height is the node's current height
func (n *node) insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isFuzz: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(path, parts, height+1) // recursive calls
}

// search if the parts have a path, return the final node
// * can only use in the end
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		} // failure in the middle of starting a business
		return n
	}

	part := parts[height]
	children := n.matchChildren(part) // Children may include Fuzz nodes

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil // no path for the parts
}

func (n *node) traverse(slice *[]*node) {
	if n.path != "" {
		*slice = append(*slice, n)
	}
	for _, child := range n.children {
		child.traverse(slice)
	}
}
