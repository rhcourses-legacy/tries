package tries

// Node is a trie node.
// It contains data and a map of children nodes.
//   - data is a slice of strings because a node's data in our examples is a list of words.
//   - children is a map of runes to nodes because we traverse the trie by reading characters from a string.
type Node struct {
	data     []string
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
