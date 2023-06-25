package tries

// Node is a trie node.
// It contains data and a map of children nodes.
//   - data is a slice of any. Usually, we will be using strings, but there is no reason to restrict the data type.
//   - children is a map of runes to nodes because we traverse the trie by reading characters from a string.
type Node struct {
	data     []any
	children map[rune]*Node
}

// NewNode creates a new node.
func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
	}
}

// IsLeaf returns true if the node is a leaf.
func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

// IsEmpty returns true if the node is empty.
func (n *Node) IsEmpty() bool {
	return len(n.data) == 0
}

// IsValid returns true if the node contains a valid children map.
func (n *Node) IsValid() bool {
	return n.children != nil
}

// GetOrCreateChild expects a rune and returns the corresponding child node.
// If the child node does not exist, it is created.
func (n *Node) GetOrCreateChild(r rune) *Node {
	if _, exists := n.children[r]; !exists {
		n.children[r] = NewNode()
	}
	return n.children[r]
}

// Insert expects a string and inserts it into the trie.
// It optinally expects data strings that will be added to the node.
// If the node already exists, the data strings are appended to the node's data.
func (n *Node) Insert(s string, data ...any) {
	current := n
	for _, r := range s {
		current = current.GetOrCreateChild(r)
	}
	current.data = append(current.data, data...)
}

// Get expects a string and returns the corresponding node.
// If the node does not exist, it returns nil.
func (n *Node) Get(s string) *Node {
	current := n
	for _, r := range s {
		if current = current.children[r]; current == nil {
			return nil
		}
	}
	return current
}
