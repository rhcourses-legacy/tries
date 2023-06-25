package tries

// Node is a trie node.
// It contains a map of children nodes.
//   - children is a map of runes to nodes because we traverse the trie by reading characters from a string.
type Node struct {
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
