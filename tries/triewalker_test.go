package tries

import "testing"

// TestTrieWalker_New tests the NewTrieWalker function.
// It creates a new trie walker and checks whether it is both at the root and at a leaf.
func TestTrieWalker_New(t *testing.T) {
	trie := NewTrie()
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
	if !walker.IsAtLeaf() {
		t.Error("New trie walker should be at leaf")
	}
}
