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

// TestTrie_IsEmpty tests the IsEmpty function.
// It creates a new trie and checks whether it is empty.
// Then it adds a child node and checks whether the trie is not empty.
func TestTrie_IsEmpty(t *testing.T) {
	trie := NewTrie()
	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}
	trie.root.children['a'] = NewNode()
	if trie.IsEmpty() {
		t.Error("Trie should not be empty")
	}
}

