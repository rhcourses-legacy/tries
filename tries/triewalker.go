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
