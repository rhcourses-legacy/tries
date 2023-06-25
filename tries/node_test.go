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

