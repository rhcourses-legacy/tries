package tries

// Trie is a data structure that stores strings.
// - It contains a root node of type Node.
// - It provides methods to insert strings with data and to traverse the trie using a string.
type Trie struct {
	root *Node
}

// NewTrie creates a new trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

// IsEmpty returns true if the trie is empty.
func (t *Trie) IsEmpty() bool {
	return t.root.IsLeaf()
}

// IsValid returns true if the trie contains a valid root node.
func (t *Trie) IsValid() bool {
	return t.root != nil && t.root.IsValid()
}
