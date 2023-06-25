package tries

import "testing"

// TestTrieWalker_New_Empty tests the NewTrieWalker function.
// It creates a trie walker with an empty trie and checks whether it is both at the root and at a leaf.
func TestTrieWalker_New_Empty(t *testing.T) {
	trie := NewTrie()
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
	if !walker.IsAtLeaf() {
		t.Error("New trie walker should be at leaf")
	}
}

// TestTrieWalker_New_nonempty tests the NewTrieWalker function.
// It creates a trie walker and inserts a node into the trie.
// It checks whether the walker is at the root and not at a leaf.
func TestTrieWalker_New_nonempty(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
	if walker.IsAtLeaf() {
		t.Error("New trie walker should not be at leaf")
	}
}
