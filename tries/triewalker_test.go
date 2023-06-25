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

// TestTrieWalker_Step tests the Step function.
// It creates a new trie walker, adds a node to the trie, and steps past it.
// It checks whether the following:
//   - the Step function returns true up to the node
//   - the Step function returns false past the node
//   - each step moves the walker to the correct node
func TestTrieWalker_Step(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)

	a := trie.root.Get("a")
	b := trie.root.Get("ab")
	c := trie.root.Get("abc")

	if !walker.Step('a') {
		t.Error("Step should return true for rune 'a'")
	}
	if walker.current != a {
		t.Error("Step should move walker to correct node")
	}

	if !walker.Step('b') {
		t.Error("Step should return true for rune 'b'")
	}
	if walker.current != b {
		t.Error("Step should move walker to correct node")
	}

	if !walker.Step('c') {
		t.Error("Step should return true for rune 'c'")
	}

	if walker.Step('d') {
		t.Error("Step should return false for rune 'd'")
	}
	if walker.current != c {
		t.Error("Step should not move walker past leaf node")
	}
}

// TestTrieWalker_Walk_match tests the Walk function.
// It creates a new trie walker, adds a node to the trie, and walks a string that matches the node.
// It checks the following:
//   - The Walk function returns the correct number of consumed runes.
//   - The Walk function moves the walker to the correct node.
func TestTrieWalker_Walk_match(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)

	if n := walker.Walk("abc"); n != 3 {
		t.Errorf("Walk should consume 3 runes, but consumed %d", n)
	}
	if walker.current != trie.root.Get("abc") {
		t.Error("Walk should move walker to correct node")
	}
}

// TestTrieWalker_Walk_walkpast tests the Walk function.
// It creates a new trie walker, adds a node to the trie, and walks a string that
// hits a node but would walk past it.
// It checks the following:
//   - The Walk function returns the correct number of consumed runes.
//   - The Walk function moves the walker to the inserted node.
func TestTrieWalker_Walk_walkpast(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)

	if n := walker.Walk("abcd"); n != 3 {
		t.Errorf("Walk should consume 3 runes, but consumed %d", n)
	}
	if walker.current != trie.root.Get("abc") {
		t.Error("Walk should move walker to correct node")
	}
}

// TestTrieWalker_Walk_short tests the Walk function.
// It creates a new trie walker, adds a node to the trie, and walks a string that
// is shorter than the one corresponding to the node.
// It checks the following:
//   - The Walk function returns the correct number of consumed runes.
//   - The Walk function moves the walker to the correct inner node.
func TestTrieWalker_Walk_short(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)

	if n := walker.Walk("ab"); n != 2 {
		t.Errorf("Walk should consume 2 runes, but consumed %d", n)
	}
	if walker.current != trie.root.Get("ab") {
		t.Error("Walk should move walker to correct node")
	}
}
