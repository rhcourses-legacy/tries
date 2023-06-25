package simplezip

import "github.com/rhcourses/tries/tries"

// Unzipper is a data structure that stores a compressed string and everything that is
// needed to track the progress while decompressing it.

type Unzipper struct {
	compressed string

	tw       *tries.TrieWalker
	position int

	result string
}

// NewUnzipper creates a new unzipper from a compressed string and a trie.
func NewUnzipper(compressed string, trie *tries.Trie) *Unzipper {
	return &Unzipper{
		tw:         tries.NewTrieWalker(trie),
		compressed: compressed,
	}
}

// Done returns true if the unzipper is done.
func (uz *Unzipper) Done() bool {
	return uz.position == len(uz.compressed)
}

// Result returns the decompressed string.
func (uz *Unzipper) Result() string {
	return uz.result
}

// Run decompresses the string.
func (uz *Unzipper) Run() {
	for !uz.Done() {
		consumed := uz.tw.Walk(uz.compressed[uz.position:])
		if consumed == 0 {
			uz.result += string(uz.compressed[uz.position])
			uz.position++
			continue
		}
		data := uz.tw.Data()
		if len(data) != 0 {
			uz.result += data[0].(string)
		} else {
			uz.result += uz.compressed[uz.position : uz.position+consumed]
		}
		uz.tw.Reset()
		uz.position += consumed
	}
}
