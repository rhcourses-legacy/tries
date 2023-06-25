package tries

// TrieWalker is a struct that is used to traverse a trie.
// It contains a trie and a pointer to the current node.
// It provides methods that allow to traverse the trie by consuming strings or runes.
type TrieWalker struct {
	trie    *Trie
	current *Node
}

// NewTrieWalker creates a new trie walker.
func NewTrieWalker(trie *Trie) *TrieWalker {
	return &TrieWalker{
		trie:    trie,
		current: trie.root,
	}
}

// IsAtRoot returns true if the walker is at the root node.
func (tw *TrieWalker) IsAtRoot() bool {
	return tw.current == tw.trie.root
}

// Reset moves the walker to the root node.
func (tw *TrieWalker) Reset() {
	tw.current = tw.trie.root
}

// IsAtLeaf returns true if the walker is at a leaf node.
func (tw *TrieWalker) IsAtLeaf() bool {
	return tw.current.IsLeaf()
}

// Step expects a rune and moves the walker to the corresponding child node.
// If the child node does not exist, it returns false and does not move the walker.
// Returns true if the step was successful.
func (tw *TrieWalker) Step(r rune) bool {
	next := tw.current.children[r]
	if next == nil {
		return false
	}
	tw.current = next
	return true
}

// Walk expects a string and moves the walker until either of the following:
//   - the walker is at a leaf node
//   - the walker is at the end of the string
//
// Returns the number of runes that were consumed.
func (tw *TrieWalker) Walk(s string) int {
	for i, r := range s {
		if !tw.Step(r) {
			return i
		}
	}
	return len(s)
}

// Data returns the data of the current node.
func (tw *TrieWalker) Data() []any {
	return tw.current.data
}
