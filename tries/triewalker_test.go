package tries

import "testing"

// TestTrieWalker_New tests the NewTrieWalker function.
// It creates a new trie walker and checks whether it is at the root.
func TestTrieWalker_New(t *testing.T) {
	trie := NewTrie()
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
}
