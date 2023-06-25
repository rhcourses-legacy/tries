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

// Insert expects a string and inserts it into the trie.
// It optinally expects data strings that will be added to the node.
// If the node already exists, the data strings are appended to the node's data.
func (t *Trie) Insert(s string, data ...any) {
	t.root.Insert(s, data...)
}

// GetData expects a string and returns the corresponding data.
// If the node does not exist, it returns nil.
func (t *Trie) GetData(s string) []any {
	if node := t.root.Get(s); node != nil {
		return node.data
	}
	return nil
}
