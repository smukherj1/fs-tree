package fstree

import (
	"testing"
)

func TestEmptyTree(t *testing.T) {
	n := NewTree()
	if n.Name != "" {
		t.Errorf("Root node of new empty tree had unexpected name: %q", n.Name)
	}
	if n.Children != nil {
		t.Errorf("Root node of new empty tree had unexpected value for field Children, want nil, got: list of len %d", len(n.Children))
	}
	if n.Parent != nil {
		t.Error("Root node of new empty tree had unexpected value for field Parent. want nil, got not nil")
	}
	if n.Path() != "" {
		t.Errorf("Root node of new empty tree returned unexpected value for path, want \"\", got %q", n.Path())
	}
	if c := n.Find("foo"); c != nil {
		t.Errorf("Root of node of new empty tree unexpectedly returned node for path \"foo\", want nil, got %s", c.Path())
	}
}

func TestNodeCreate(t *testing.T) {
	n := NewTree()
	c := n.Get("foo")
	if c == nil {
		t.Fatal("Failed to create node at path \"foo\" in new tree")
	}
	if c.Name != "foo" {
		t.Errorf("Unexpected name for newly created node, want \"foo\", got %q", c.Name)
	}
	if p := c.Path(); p != "foo" {
		t.Errorf("Unexpected path returned for newly created node, want \"foo\", got %q", p)
	}
}
