package tries

import "testing"

// TestNode_New_IsLeaf tests the NewNode function.
// It creates a new node and checks whether it is a leaf.
func TestNode_New_IsLeaf(t *testing.T) {
	n := NewNode()
	if !n.IsLeaf() {
		t.Error("New node should be a leaf")
	}
}

// TestNode_New_IsEmpty tests the NewNode function.
// It creates a new node and checks whether it is empty.
func TestNode_New_IsEmpty(t *testing.T) {
	n := NewNode()
	if !n.IsEmpty() {
		t.Error("New node should be empty")
	}
}

// TestNode_New_IsValid tests the NewNode function.
// It creates a new node and checks whether it is valid.
func TestNode_New_IsValid(t *testing.T) {
	n := NewNode()
	if !n.IsValid() {
		t.Error("New node should be valid")
	}
}

// TestNode_GetOrCreateChild tests the GetOrCreateChild function.
// It creates a new node and uses GetOrCreateChild to create child nodes.
//   - During creation, it checks whether each created node is a valid empty leaf.
//   - After all child nodes are created, it checks whether the last node is still a leaf,
//     but the others are not. All nodes must still be valid and empty.
func TestNode_GetOrCreateChild(t *testing.T) {
	root := NewNode()

	current := root
	for _, r := range "abc" {
		child := current.GetOrCreateChild(r)
		if child == nil {
			t.Errorf("Child node for rune %v should be created", r)
		}
		if !child.IsValid() {
			t.Errorf("Child node for rune %v should be valid", r)
		}
		if !child.IsEmpty() {
			t.Errorf("Child node for rune %v should be empty", r)
		}
		current = child
	}

	a := root.children['a']
	b := a.children['b']
	c := b.children['c']

	// Check validity and emptiness.
	for _, node := range []*Node{a, b, c} {
		if !node.IsValid() {
			t.Errorf("Node %v should be valid", node)
		}
		if !node.IsEmpty() {
			t.Errorf("Node %v should be empty", node)
		}
	}

	// Check leaf property.
	for _, node := range []*Node{a, b} {
		if node.IsLeaf() {
			t.Errorf("Node %v should not be a leaf", node)
		}
	}
	if !c.IsLeaf() {
		t.Errorf("Node %v should be a leaf", c)
	}
}
