package fstree

import (
	"strings"
)

// Node represents an item in the filesystem like tree.
type Node struct {
	// Name is the name of this node. Every node other than the root node can have a name.
	Name string
	// Parent is the parent node of this node. This field is trivially nil for the root node.
	Parent *Node
	// Children are the children nodes of this node.
	Children []*Node
	// Payload can contain any object relevant to this node in the tree.
	Payload interface{}
}

// NewTree creates and returns a new root node.
func NewTree() *Node {
	return &Node{}
}

func (n *Node) findChild(name string) *Node {
	for _, c := range n.Children {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// Find attempts to traverse the tree rooted at the given node (where the given node need not be an actual root node in a tree) to find the node represented by the given path.
func (n *Node) Find(nodePath string) *Node {
	if nodePath == "" || nodePath == "/" {
		return n
	}
	splitPath := strings.Split(nodePath, "/")
	for i := 0; i < len(splitPath); i++ {
		c := n.findChild(splitPath[i])
		if c == nil {
			return nil
		}
		n = c
	}
	return n
}

// Get returns the node at the given path. If such a node doesn't exist, every node along the path to the requested node is created and the final node is returned.
func (n *Node) Get(nodePath string) *Node {
	if nodePath == "" || nodePath == "/" {
		return n
	}
	splitPath := strings.Split(nodePath, "/")
	for i := 0; i < len(splitPath); i++ {
		c := n.findChild(splitPath[i])
		if c == nil {
			c = &Node{
				Name:   splitPath[i],
				Parent: n,
			}
		}
		n = c
	}
	return n
}

// Path returns the path to the given node from the root of the tree its in.
func (n *Node) Path() string {
	p := []string{}
	for n != nil && n.Parent != nil {
		p = append(p, n.Name)
		n = n.Parent
	}
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return strings.Join(p, "/")
}
