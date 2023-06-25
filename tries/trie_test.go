package tries

import "testing"

// TestTrie_New tests the NewTrie function.
// It creates a new trie and checks whether it is empty and valid.
func TestTrie_New(t *testing.T) {
	trie := NewTrie()
	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}
	if !trie.IsValid() {
		t.Error("New trie should be valid")
	}
}
